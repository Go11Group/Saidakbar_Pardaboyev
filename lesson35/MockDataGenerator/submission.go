package mockdatagenerator

import (
	"fmt"
	"leetcode/model"
	"leetcode/storage/postgres"
	"math/rand"
)

func SubmissionGenerator(Submissions *postgres.Submissions, users []model.User,
	problems []model.Problem, languages []model.Language) error {
	code := "int* twoSum(int* nums, int numsSize, int target, int* returnSize){\n\tint* result = (int*)malloc(2 * sizeof(int));\n*returnSize = 2;\n\nfor(int i = 0; i < numsSize - 1; i++) {\n\tfor(int j = i + 1; j < numsSize; j++) {\n\t\tif(nums[i] + nums[j] == target) {\n\t\t\tresult[0] = i;\n\t\t\tresult[1] = j;\n\t\t\treturn result;\n\t\t}\n\t}\n}\n\nreturn result;\n}"
	submissionStatuses := []string{"Eccepted", "Wrong Answer", "Time Limit Exceeded", "Memory Limit Exceeded", "Runtime Error", "Compile Error"}

	for userIndex, user := range users {
		usertId := user.Id
		numberOfProblemsUserSubmitted := rand.Intn(len(problems))
		for i := 0; i < numberOfProblemsUserSubmitted; i++ {
			problemIndex := rand.Intn(len(problems))
			problemId := problems[problemIndex].Id

			languageIdnex := rand.Intn(len(languages))
			language := languages[languageIdnex].Id

			numberOfSubmissionForQuestion := rand.Intn(10)
			for j := 0; j < numberOfSubmissionForQuestion; j++ {
				submissionStatusIndex := rand.Intn(len(submissionStatuses))
				submissionStatus := submissionStatuses[submissionStatusIndex]

				newSubmission := model.Submission{
					UserId:           usertId,
					ProblemId:        problemId,
					Language:         language,
					SubmissionStatus: submissionStatus,
					Code:             code,
				}
				err := Submissions.CreateSubmission(&newSubmission)
				if err != nil {
					return err
				}
			}
		}
		if userIndex%10 == 0 && userIndex != 0 {
			fmt.Printf("%d ta odam uchun mock submission qo'shdi\n", userIndex)
		}
	}

	return nil
}
