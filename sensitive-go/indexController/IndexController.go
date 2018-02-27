package indexController

import (
    "fmt"
    "net/http"
    "gopkg.in/mgo.v2"

    "encoding/json"

    "github.com/julienschmidt/httprouter"

    "sensitive-go/sensitive-go/dao"
    "sensitive-go/sensitive-go/model"
    "sensitive-go/sensitive-go/wordFilter"
)

type Controller struct{
    session *mgo.Session
}

type ResultData struct {
    Success bool        `json:"success"`
    Code string         `json:"code"`      
    Message string      `json:"message"`
    Word string         `json:"word"`
}
/**
 * add sensitive controller function
 */
func (clt *Controller) addSensitive (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    var result ResultData

    dao := dao.Dao{
        S:clt.session,
    }

    r.ParseForm()
    requestParams := r.Form
    word := requestParams["word"][0]

    WordStruct := &model.WordStruct{
        Word:word,
    }

    queryResult  := dao.FindOne(word)
    if  queryResult {
        result.Code = "501"
        result.Success = false
        result.Message = "this sensitive word is having"
    }else{
        // orm
        err := dao.Insert(WordStruct)

        var wordlist []string
        wordlist = append(wordlist,params.ByName("word"))
        wordFilter.LoadSensitiveWord(wordlist)
        if err == nil {
            fmt.Println("add sensitive success")
            result.Code = "200"
            result.Success = true
            result.Message = "success"
        }else{
            fmt.Println("add sensitive false")
            result.Code = "500"
            result.Success = false
            result.Message = "false"
        }
        fmt.Println(result)
    }
    jsondata , _ := json.Marshal(result)

    fmt.Fprint(w, string(jsondata))
}

/**
 * add sensitive controller function
 */
func (clt *Controller) delSensitive (w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    var result ResultData

    dao := dao.Dao{
        S:clt.session,
    }

    r.ParseForm()
    requestParams := r.Form
    word := requestParams["word"][0]

    err := dao.Delete(word)

    wordFilter.DelSensitiveWord(word)

    if err == nil {
        fmt.Println("del sensitive true")
        result.Code = "200"
        result.Success = true
        result.Message = "success"
    }else{
        fmt.Println("del sensitive false")
        result.Code = "500"
        result.Success = false
        result.Message = "false"
    }
    jsondata , _ := json.Marshal(result)

    fmt.Fprint(w, string(jsondata))
}

/**
 *  replace sensitive word to '*' 
 */
func (clt *Controller) check (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    var result ResultData

    r.ParseForm()
    requestParams := r.Form
    text := requestParams["txt"][0]

    resultText := wordFilter.ReplaceSensitiveWord(text)

    result.Code = "200"
    result.Success = true
    result.Message = "success"
    result.Word = resultText

    data ,jsonerr := json.Marshal(result)
    if jsonerr != nil {
        fmt.Println("check sensitive false ...")
    }
    fmt.Fprint(w, string(data))
}


