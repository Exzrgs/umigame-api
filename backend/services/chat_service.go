package services

import (
	"context"
	"encoding/json"
	"os"
	"umigame-api/models"
	"umigame-api/repositories"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func (s *Service) GetChatroomListService() ([]models.Chatroom, error) {
	chatrooms, err := repositories.SelectChatroomList(s.db, s.userID)
	if err != nil {
		return nil, err
	}

	return chatrooms, nil
}

func (s *Service) PostQuestionService(chat models.Chat) (models.Chat, error) {
	problemBase, err := repositories.SelectProblemBase(s.db, chat.ProblemID)
	if err != nil {
		return models.Chat{}, err
	}

	chat.UserID = s.userID

	answerChat, err := attachAnswer(problemBase, chat)
	if err != nil {
		return models.Chat{}, err
	}

	if err := repositories.InsertChat(s.db, answerChat); err != nil {
		return models.Chat{}, err
	}

	return answerChat, nil
}

func attachAnswer(problem models.Problem, chat models.Chat) (models.Chat, error) {
	message := createUserMessage(problem, chat)

	answerJson, err := getGPTResponse(message)
	if err != nil {
		return models.Chat{}, err
	}

	json.Unmarshal([]byte(answerJson), &chat)

	return chat, nil
}

func createUserMessage(problem models.Problem, chat models.Chat) string {
	message := `あなたは水平思考ゲームのジャッジをします。問題文とその答えが与えられるので、質問に対して、「はい」なら1、「いいえ」なら0、「どちらともいえない、答えにくい」なら2のいずれかで回答してください。` +
		"問題文: " +
		problem.Statement +
		"答え: " +
		problem.Answer +
		"質問文: " +
		chat.Question

	return message
}

func getGPTResponse(message string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"answer": {
				Type:        jsonschema.Integer,
				Description: "number of answer",
			},
		},
		Required: []string{"answer"},
	}
	f := openai.FunctionDefinition{
		Name:        "get_answer",
		Description: "Get answer to question for problem",
		Parameters:  params,
	}
	t := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: f,
	}
	requestMessage := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: requestMessage,
			Tools:    []openai.Tool{t},
		},
	)

	if err != nil {
		return "", err
	}

	respJson := resp.Choices[0].Message.ToolCalls[0].Function.Arguments

	return respJson, nil
}
