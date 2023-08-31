package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-operations-doc-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-operations-doc-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var headerDoc *[]dpfm_api_output_formatter.HeaderDoc
	var itemDoc *[]dpfm_api_output_formatter.ItemDoc

	for _, fn := range accepter {
		switch fn {
		case "HeaderDoc":
			func() {
				headerDoc = c.HeaderDoc(input, output, errs, log)
			}()
		case "ItemDoc":
			func() {
				itemDoc = c.ItemDoc(input, output, errs, log)
			}()
		}
	}

	data := &dpfm_api_output_formatter.Message{
		HeaderDoc: headerDoc,
		ItemDoc:   itemDoc,
	}

	return data
}

func (c *DPFMAPICaller) HeaderDoc(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.HeaderDoc {
	where := "WHERE 1 = 1"

	if input.HeaderDoc.Operations != nil {
		where = fmt.Sprintf("%s\nAND Operations = %d", where, *input.HeaderDoc.Operations)
	}
	if input.HeaderDoc.DocType != nil && len(*input.HeaderDoc.DocType) != 0 {
		where = fmt.Sprintf("%s\nAND DocType = '%v'", where, *input.HeaderDoc.DocType)
	}
	if input.HeaderDoc.DocIssuerBusinessPartner != nil && *input.HeaderDoc.DocIssuerBusinessPartner != 0 {
		where = fmt.Sprintf("%s\nAND DocIssuerBusinessPartner = %v", where, *input.HeaderDoc.DocIssuerBusinessPartner)
	}
	groupBy := "\nGROUP BY Operations, DocType, DocIssuerBusinessPartner "

	rows, err := c.db.Query(
		`SELECT
    Operations, DocType, MAX(DocVersionID), DocID, FileExtension, FileName, FilePath, DocIssuerBusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_header_doc_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeaderDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemDoc(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemDoc {
	where := "WHERE 1 = 1"

	if input.HeaderDoc.Operations != nil {
		where = fmt.Sprintf("%s\nAND Operations = %d", where, *input.HeaderDoc.Operations)
	}
	if input.HeaderDoc.ItemDoc.OperationsItem != nil {
		where = fmt.Sprintf("%s\nAND OperationsItem = %d", where, *input.HeaderDoc.ItemDoc.OperationsItem)
	}
	if input.HeaderDoc.ItemDoc.DocType != nil {
		where = fmt.Sprintf("%s\nAND DocType = '%v'", where, *input.HeaderDoc.ItemDoc.DocType)
	}
	if input.HeaderDoc.ItemDoc.DocIssuerBusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND DocIssuerBusinessPartner = %v", where, *input.HeaderDoc.ItemDoc.DocIssuerBusinessPartner)
	}
	groupBy := "\nGROUP BY Operations, OperationsItem, DocType, DocIssuerBusinessPartner "

	rows, err := c.db.Query(
		`SELECT
    Operations, OperationsItem, DocType, MAX(DocVersionID), DocID, FileExtension, FileName, FilePath, DocIssuerBusinessPartner
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_operations_item_doc_data
		` + where + groupBy + `;`)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItemDoc(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
