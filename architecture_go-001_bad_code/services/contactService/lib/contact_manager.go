package lib

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"forgoproject/services/contactService/proto"
)

type ContactManager struct {
	proto.UnimplementedContactServiceServer
	DB *sql.DB
}

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func NewContactManager(DB *sql.DB) *ContactManager {
	return &ContactManager{
		DB: DB,
	}
}

func (c *ContactManager) createWithId(ctx context.Context, in *proto.Contact) (*proto.Contact, error) {
	const errContactAlreadyExists = "pq: duplicate key value violates unique constraint \"pk_contact\""

	// Работа с Атрибутами
	userAttrLists, err := c.GetAttributeKeyList(ctx, &proto.AttributeOwnerId{Id: in.CreatedBy})
	if err != nil {
		return nil, err
	}

	in.Attributes, err = checkContactAttributes(in.Attributes, userAttrLists)
	if err != nil {
		return nil, err
	}

	attrBytes, err := json.Marshal(in.Attributes)
	if err != nil {
		return nil, err
	}

	phoneNumber := convertPhoneNumberToFormatWithoutLength(in.PhoneNumber)

	var gtm sql.NullString

	if in.Gtm == "" {
		// timezoneByNumber := c.gtmHelper.GetTimezoneByNumber(ctx, phoneNumber)
		// gtm = sql.NullString{
		// 	String: timezoneByNumber,
		// 	Valid:  !strings.EqualFold(timezoneByNumber, ""),
		// }
	} else {
		gtm = sql.NullString{
			String: in.Gtm,
			Valid:  !strings.EqualFold(in.Gtm, ""),
		}
	}

	// Validate phoneNumber
	if len(phoneNumber) == 0 {
		return nil, err
	}

	query, args, err := psql.
		Insert("clients.contact").
		SetMap(map[string]interface{}{
			"id":           in.Id,
			"full_name":    in.FullName,
			"email":        in.Email,
			"phone_number": phoneNumber,
			"age":          in.Age,
			"gender":       in.Gender,
			"city":         in.City,
			"file_name":    in.FileName,
			"created_at":   time.Now().UTC(),
			"modified_at":  time.Now().UTC(),
			"created_by":   in.CreatedBy,
			"attributes":   string(attrBytes),
			"name":         in.Name,
			"surname":      in.Surname,
			"patronymic":   in.Patronymic,
			"parameter_1":  in.Parameter_1,
			"parameter_2":  in.Parameter_2,
			"gtm":          gtm,
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = c.DB.ExecContext(ctx, query, args...)
	if err != nil {
		if err.Error() == errContactAlreadyExists {
			return nil, err
		}

		return nil, err
	}

	return &proto.Contact{Id: in.Id}, nil
}

func (c *ContactManager) Create(ctx context.Context, in *proto.Contact) (*proto.Contact, error) {
	id := uuid.New()

	in.Id = id.String()

	return c.createWithId(ctx, in)
}

func (c *ContactManager) Update(ctx context.Context, in *proto.Contact) (*proto.Contact, error) {
	// Работа с Атрибутами
	userAttrLists, err := c.GetAttributeKeyList(ctx, &proto.AttributeOwnerId{Id: in.CreatedBy})
	if err != nil {
		return nil, err
	}

	in.Attributes, err = checkContactAttributes(in.Attributes, userAttrLists)
	if err != nil {
		return nil, err
	}

	attrBytes, err := json.Marshal(in.Attributes)
	if err != nil {
		return nil, err
	}
	phone := sql.NullString{
		String: convertPhoneNumberToFormatWithoutLength(in.PhoneNumber),
		Valid:  !strings.EqualFold(in.PhoneNumber, ""),
	}

	var gtm sql.NullString
	if in.Gtm == "" {
		// timezoneByNumber := c.gtmHelper.GetTimezoneByNumber(ctx, phone.String)
		// gtm = sql.NullString{
		// 	String: timezoneByNumber,
		// 	Valid:  !strings.EqualFold(timezoneByNumber, ""),
		// }
	} else {
		gtm = sql.NullString{
			String: in.Gtm,
			Valid:  !strings.EqualFold(in.Gtm, ""),
		}
	}

	// Validate phoneNumber
	if len(in.GetPhoneNumber()) == 0 {
		return nil, err
	}

	query, args, err := psql.
		Update("clients.contact").
		Set("full_name", in.FullName).
		Set("email", in.Email).
		Set("phone_number", phone).
		Set("age", in.Age).
		Set("gender", in.Gender).
		Set("city", in.City).
		Set("modified_at", time.Now().UTC()).
		Set("attributes", string(attrBytes)).
		Set("name", in.Name).
		Set("surname", in.Surname).
		Set("patronymic", in.Patronymic).
		Set("parameter_1", in.Parameter_1).
		Set("parameter_2", in.Parameter_2).
		Set("gtm", gtm).
		Where(squirrel.Eq{
			"id":          in.Id,
			"created_by":  in.CreatedBy,
			"is_archived": false,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	res, errExec := c.DB.ExecContext(ctx, query, args...)
	if errExec != nil {
		return nil, err
	}

	rowsAffected, errRows := res.RowsAffected()
	if errRows != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return c.createWithId(ctx, in)
	}

	return in, nil
}

func (c *ContactManager) ReadById(ctx context.Context, in *proto.ContactIdAndCreatedBy) (*proto.Contact, error) {
	sqlQuery := `
SELECT c.id,
	c.age,
	c.email,
	c.file_name,
	c.full_name,
	c.gender,
	c.phone_number,
	c.is_blacklist,
    array_to_string(array_agg(cInG.group_id), ',') AS inGroups,
	COALESCE((c.attributes)::text, '') as attributes,
	COALESCE((c.gtm)::text, '') as gtm,
	c.name,
	c.surname,
	c.patronymic,
	c.parameter_1,
	c.parameter_2
FROM clients.contact as c
-- Ищу группы, в которых есть контакт
LEFT OUTER JOIN clients.contact_in_group as cInG ON cInG.contact_id = c.id
WHERE c.id = $1
    AND c.created_by = $2
    AND c.is_archived = FALSE
GROUP BY c.id
`
	row := c.DB.QueryRow(sqlQuery, in.Id, in.UserUuid)
	contact := proto.Contact{
		Attributes: &proto.Contact_AttributeList{
			List: []*proto.Contact_AttributeList_AttributeItem{},
		},
	}

	var attributes string

	err := row.Scan(
		&contact.Id,
		&contact.Age,
		&contact.Email,
		&contact.FileName,
		&contact.FullName,
		&contact.Gender,
		&contact.PhoneNumber,
		&contact.IsBlacklist,
		&contact.InGroups,
		&attributes,
		&contact.Gtm,
		&contact.Name,
		&contact.Surname,
		&contact.Patronymic,
		&contact.Parameter_1,
		&contact.Parameter_2,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return &contact, nil
		default:
			return nil, err
		}
	}

	err = json.Unmarshal([]byte(attributes), contact.Attributes)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (c *ContactManager) GetAttributeKeyList(ctx context.Context, in *proto.AttributeOwnerId) (*proto.ContactAttributeKeyList, error) {
	attributeKeys := proto.ContactAttributeKeyList{
		Total: 0,
		List:  []*proto.ContactAttributeKey{},
	}

	// ---------
	// TOTAL
	// ---------

	// Считаем количество всех атрибутов
	sqlQueryTotal := `
		SELECT count(a.id)
		FROM clients.contact_attribute as a
		WHERE a.created_by = $1
		AND a.is_archived = FALSE
`

	rowTotal := c.DB.QueryRow(sqlQueryTotal, in.Id)
	var total int32

	errTotal := rowTotal.Scan(&total)
	if errTotal != nil && errTotal != sql.ErrNoRows {
		return nil, errTotal
	}

	if total == 0 {
		return &attributeKeys, nil
	}

	attributeKeys.Total = total

	// ---------
	// LIST
	// ---------

	sqlQuery := `
		SELECT a.id, a.key_title
		FROM clients.contact_attribute as a
		WHERE a.created_by = $1
		AND a.is_archived = FALSE
		ORDER BY a.key_title
`

	rows, err := c.DB.Query(sqlQuery, in.Id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			attributeKeys.Total = 0

			return &attributeKeys, nil
		default:
			return nil, err
		}
	}
	defer rows.Close()

	for rows.Next() {
		attributeKey := &proto.ContactAttributeKey{}

		err = rows.Scan(
			&attributeKey.Id,
			&attributeKey.KeyTitle,
		)
		if err != nil {
			return nil, err
		}

		attributeKeys.List = append(attributeKeys.List, attributeKey)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &attributeKeys, nil
}

// Проверяем аттрибуты контакта
func checkContactAttributes(attributes *proto.Contact_AttributeList, userAttrList *proto.ContactAttributeKeyList) (attributesSuccessList *proto.Contact_AttributeList, err error) {
	// Собираем только успешные
	attributesSuccessList = &proto.Contact_AttributeList{
		List: []*proto.Contact_AttributeList_AttributeItem{},
	}
	// Проверяем id присланных кастомных полей
	attrIds := []string{}
	if attributes != nil && attributes.List != nil && len(attributes.List) > 0 {
		for _, a := range attributes.List {
			if _, err := uuid.Parse(a.Id); err != nil {
				return nil, err
			}
			attrIds = append(attrIds, a.Id)
		}

		for _, userAttr := range userAttrList.List {
			for _, attr := range attributes.List {
				if userAttr.Id == attr.Id {
					attributesSuccessList.List = append(attributesSuccessList.List, attr)
				}
			}
		}
	}
	return attributesSuccessList, nil
}

func getNumbers(input string) string {
	var number string
	for _, t := range input {
		if t >= 48 && t <= 57 { // 48 - 57 in ASCII this numbers   0-9
			number += string(t)
		}
	}
	return number
}

func convertPhoneNumberToFormatWithoutLength(input string) (res string) {
	input = getNumbers(input)

	if strings.HasPrefix(input, "+") {
		input = strings.Replace(input, "+", "", 1)
	}

	// Если номер в таком формате: 81234567890, то заменяем первую 8 на 7 (Россия)
	if strings.HasPrefix(input, "8") && len(input) == 11 {
		input = strings.Replace(input, "8", "7", 1)
		// Если номер в таком формате: 1234567890, то добавляем первой цифорой 7 (Россия)
	}

	res = input

	return
}
