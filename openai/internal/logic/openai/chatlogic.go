package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-zero-chatgpt/common/ctxdata"
	"go-zero-chatgpt/openai/internal/svc"
	"go-zero-chatgpt/openai/internal/types"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	resp = ""
	UserId := ctxdata.GetUidFromCtx(l.ctx)

	ApiUser, err := l.svcCtx.OpenaiModel.FindOne(l.ctx, UserId)
	if err != nil {
		return resp, err
	}
	Api := ApiUser.Api

	resp, err = OpenAiChat(Api, req.Question)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func OpenAiChat(apiKey string, question string) (result string, err error) {

	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	result = ""
	// 构建请求体
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":      question,
		"max_tokens":  5,
		"temperature": 0.5,
		"n":           1,
	})
	if err != nil {
		return result, err
	}

	// 发送 HTTP 请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// 解析响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return result, err
	}

	// 输出结果
	var builder strings.Builder
	for _, choice := range response["choices"].([]interface{}) {
		text := choice.(map[string]interface{})["text"].(string)
		builder.WriteString(text)
	}
	result = builder.String()
	return result, nil
}
