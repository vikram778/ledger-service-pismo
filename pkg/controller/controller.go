package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io"
	"ledger-service-pismo/pkg/errs"
	"ledger-service-pismo/pkg/log"
	"ledger-service-pismo/pkg/models"
	service "ledger-service-pismo/pkg/service"
	"net/http"
	"strconv"
)

const (
	CT     = "Content-Type"
	CTJson = "application/json"
)

type LedgerController struct {
	ledgerService *service.LedgerService
}

func NewLedgerController(ledgerService *service.LedgerService) *LedgerController {
	return &LedgerController{
		ledgerService: ledgerService,
	}
}

func (h *LedgerController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		account models.Account
		err     error
	)
	err = h.GetParams(&account, r)
	if err != nil {
		h.FormatException(w, err)
		return
	}
	status, err := h.ledgerService.CreateAccount(context.Background(), account)
	if err != nil {
		h.FormatException(w, err)
		return
	}
	h.JSON(w, http.StatusOK, status)
}

func (h *LedgerController) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["account_id"]

	acctId, _ := strconv.Atoi(id)
	account, err := h.ledgerService.GetAccountByID(context.Background(), int64(acctId))
	if err != nil {
		h.FormatException(w, err)
		return
	}
	h.JSON(w, http.StatusOK, account)
}

func (h *LedgerController) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var (
		transaction models.Transaction
		err         error
	)
	err = h.GetParams(&transaction, r)
	if err != nil {
		h.FormatException(w, err)
		return
	}
	status, err := h.ledgerService.CreateTransaction(context.Background(), transaction)
	if err != nil {
		h.FormatException(w, err)
		return
	}
	h.JSON(w, http.StatusOK, status)
}

// GetParams unmarshalls request body to required struct
// Param - interface object in which the request body is to be unmarshalled
// Param - http request object from which the request is to be unmarshalled to desired object
func (h *LedgerController) GetParams(o interface{}, request *http.Request) (err error) {
	ct := getContentType(request)
	if ct != CTJson {
		return errors.New("unsupported media type")
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return errors.New("error reading request body")
	}

	// Restore the io.ReadCloser to its original state
	request.Body = io.NopCloser(bytes.NewBuffer(body))

	if len(body) < 1 {
		return errs.ErrorEmptyBodyContent
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		return errs.ErrorRequestBodyInvalid
	}
	return
}

// GetContentType ...
func getContentType(req *http.Request) (ct string) {
	ct = req.Header.Get(CT)
	return
}

// FormatException - formats the application exception and returns error response
func (h *LedgerController) FormatException(r http.ResponseWriter, err error) {
	h.JSON(r, http.StatusBadRequest, errs.FormatErrorResponse(err))
}

// JSON sends a JSON response body
func (h *LedgerController) JSON(r http.ResponseWriter, code int, content interface{}) {
	if fmt.Sprint(content) == "[]" {
		emptyResponse, _ := json.Marshal(make([]int64, 0))
		Output(r, code, CTJson, emptyResponse)
		return
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.Encode(content)
	Output(r, code, CTJson, b.Bytes())
}

// Output sets a full HTTP output detail
func Output(r http.ResponseWriter, code int, ctype string, content []byte) {
	log.Info("Response ", zap.Any("Message", string(content)))
	r.Header().Set("Content-Type", ctype)
	r.WriteHeader(code)
	r.Write(content)
}
