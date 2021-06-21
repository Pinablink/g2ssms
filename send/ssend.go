package send

import (
		"encoding/json"
	    "io/ioutil"
	    "net/http"
	    "net/url"
	    "strings"
	    "errors"
)

type SSAcao string
type SStatus string

const(
	SendSms SSAcao = "sendsms"
	BulkSms SSAcao = "bulksms"
	StatusError SStatus = "error" 
	StatusSuccess SStatus = "success"
	PARAM_ACAO string = "acao"
	PARAM_LOGIN string = "login"
	PARAM_TOKEN string = "token"
	PARAM_NUMERO string = "numero"
	PARAM_MSG string = "msg"
	PARAM_CAMPANHA string = "campanha"
	PARAM_DATA string = "data"
	PARAM_HORA string = "hora"
	MSG_N_INF string = "Mensagem não foi informada"
	MSG_LEN_MAX_CARACTER string = "Tamanho máximo permitido da mensagem são de 160 caracteres"
	MSG_ERROR_PARSER string = "Ocorreu um erro no Parser do Retorno"
)

type SStatusResponse struct {
	Status string `json:"status"`
	Cause string  `json:"cause"`
	Id string `json:"id"`
}

// Mensagem que será enviada
type SMsg struct {
	Msg string
}

// Dados opcionais no envio do SMS
type SSendOptSMS struct {
	Campanha string
	Data     string
	Hora     string
}

// Estrutura com os dados para solicitação de envio de sms a um microserviço
type SSendSMS struct {
	UrlService string 
	Acao       SSAcao
	Login      string 
	Token      string
	Numero     string
	Opt        SSendOptSMS
	Msg        SMsg
}

// Solicita o envio de sms para o microserviço
func (s *SSendSMS) SendSMS() (string, error) {

	var strId string = ""

    if (SMsg{}) == s.Msg || len(strings.TrimSpace(s.Msg.Msg)) == 0 {
		
		return strId, errors.New(MSG_N_INF)
	
	} else {
		
		_, err := valLenMsg(s.Msg.Msg)

		if err != nil {
			return strId, err
		} else {
			strUrl := s.UrlService
			strAcao := string(s.Acao)
			strLogin := s.Login
			strToken := s.Token
			strNum := s.Numero
			strMsg := s.Msg.Msg
			
			if err != nil {
				return strId, err
			} else {
				params := &url.Values{}
				params.Add(PARAM_ACAO, strAcao)
				params.Add(PARAM_LOGIN, strLogin)
				params.Add(PARAM_TOKEN, strToken)
				params.Add(PARAM_NUMERO, strNum)
				params.Add(PARAM_MSG, strMsg)

				valOpt(s.Opt, params)
				request := strUrl + params.Encode()
				
				resp, err := http.Get(request)
				defer resp.Body.Close()
			
				if err != nil {
					return strId, err
				} else {

					strData, errParser := ioutil.ReadAll(resp.Body)

					if errParser != nil {
						return strId, errParser
					}

					var statusResponse SStatusResponse = SStatusResponse{}
					jsonUnMarshalErr := json.Unmarshal([]byte(string(strData)), &statusResponse)
					
					// :( - Não gostei disso aqui. Mas não pensei em nada melhor no momento da escrita
					if jsonUnMarshalErr != nil {
						return strId, errors.New(MSG_ERROR_PARSER)
					}
					
					if statusResponse.Status ==  string(StatusSuccess) {
						strId = statusResponse.Id
					} else if statusResponse.Status ==  string(StatusError) { 
						return strId, errors.New(statusResponse.Cause)
					}
					//
					

				}

			}

			
		}

	}

	return strId, nil
}

func valLenMsg(strMsg string) (int, error) {
	
	lenRet := len(strings.TrimSpace(strMsg))

	if lenRet > 160 {
		return lenRet, errors.New(MSG_LEN_MAX_CARACTER)
	}

	return lenRet, nil
}

func valOpt(refOpt SSendOptSMS, paramsRequest *url.Values)  {
     
	if (SSendOptSMS{}) != refOpt {

		if len(strings.TrimSpace(refOpt.Campanha)) > 0 {
			paramsRequest.Add(PARAM_CAMPANHA, strings.TrimSpace(refOpt.Campanha))
		}

		if len(strings.TrimSpace(refOpt.Data)) > 0 {
			paramsRequest.Add(PARAM_DATA, strings.TrimSpace(refOpt.Data))
		}

		if len(strings.TrimSpace(refOpt.Hora)) > 0 {
			paramsRequest.Add(PARAM_HORA, strings.TrimSpace(refOpt.Hora))
		}
	}

}
