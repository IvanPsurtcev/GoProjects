package request

type Request struct {
	ID            int64  `db:"id" json:"id"`
	ApplicantAddr string `db:"applicant_addr" json:"applicant_addr"`
	ReceiverAddr  string `db:"receiver_addr" json:"receiver_addr"`
	Status        int64  `db:"status" json:"status"` // 1 - pending, 2 - accepted, 0 - canceled
}

type CreateRequest struct {
	ApplicantAddr string `json:"applicant_addr"`
	ReceiverAddr  string `json:"receiver_addr"`
}

type AcceptRequest struct {
	ID int64 `json:"id"`
}

type DeclineRequest struct {
	ID int64 `json:"id"`
}

type GetInvitesRequest struct {
	Address string
}

type GetInvitesResponse struct {
	Requests []Request `json:"requests"`
}
