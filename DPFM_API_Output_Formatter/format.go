package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToHeaderDoc(rows *sql.Rows) (*[]HeaderDoc, error) {
	defer rows.Close()
	headerDoc := make([]HeaderDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &HeaderDoc{}

		err := rows.Scan(
			&pm.Operations,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &headerDoc, err
		}

		data := pm
		headerDoc = append(headerDoc, HeaderDoc{
			Operations:           	  data.Operations,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &headerDoc, nil
	}

	return &headerDoc, nil
}

func ConvertToItemDoc(rows *sql.Rows) (*[]ItemDoc, error) {
	defer rows.Close()
	itemDoc := make([]ItemDoc, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &ItemDoc{}

		err := rows.Scan(
			&pm.Operations,
			&pm.OperationsItem,
			&pm.DocType,
			&pm.DocVersionID,
			&pm.DocID,
			&pm.FileExtension,
			&pm.FileName,
			&pm.FilePath,
			&pm.DocIssuerBusinessPartner,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &itemDoc, err
		}

		data := pm
		itemDoc = append(itemDoc, ItemDoc{
			Operations:           	  data.Operations,
			OperationsItem:       	  data.OperationsItem,
			DocType:                  data.DocType,
			DocVersionID:             data.DocVersionID,
			DocID:                    data.DocID,
			FileExtension:            data.FileExtension,
			FileName:                 data.FileName,
			FilePath:                 data.FilePath,
			DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &itemDoc, nil
	}

	return &itemDoc, nil
}
