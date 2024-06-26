package handler

import (
	"encoding/json"
	"leetcode/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Read
func (h *Handler) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	filter := model.FilterSubmission{}
	query := r.URL.Query()
	if query.Has("problem_id") {
		problemId := query.Get("problem_id")
		filter.ProblemId = &problemId
	}
	if query.Has("user_id") {
		userId := query.Get("user_id")
		filter.UserId = &userId
	}
	if query.Has("language_id") {
		languageId := query.Get("language_id")
		filter.Language = &languageId
	}
	if query.Has("code") {
		code := query.Get("code")
		filter.Code = &code
	}
	if query.Has("submission_status") {
		submissionStatus := query.Get("submission_status")
		filter.SubmissionStatus = &submissionStatus
	}

	submissions, err := h.Submissions.GetSubmissions(&filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting Submissions", err)
		return
	}
	err = json.NewEncoder(w).Encode(submissions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding Submissions", err)
		return
	}

}

func (h *Handler) GetSubmissionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	submission, err := h.Submissions.GetSubmissionById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error getting Submission by Id", err)
		return
	}
	err = json.NewEncoder(w).Encode(submission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while encoding Submission", err)
		return
	}
}

// Create
func (h *Handler) CreateSubmission(w http.ResponseWriter, r *http.Request) {
	newSubmission := model.Submission{}
	err := json.NewDecoder(r.Body).Decode(&newSubmission)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding Submission", err)
		return
	}
	err = h.Submissions.CreateSubmission(&newSubmission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while creating Submission", err)
		return
	}
}

// Update
func (h *Handler) UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	Submission := model.Submission{}
	vars := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&Submission)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while decoding Submission", err)
		return
	}

	Submission.Id = vars["id"]
	err = h.Submissions.UpdateSubmission(&Submission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while updating Submission", err)
		return
	}
}

// Delete
func (h *Handler) DeleteSubmission(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	err := h.Submissions.DeleteSubmission(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error while deleting Submission", err)
		return
	}
}
