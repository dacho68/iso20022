package model

// Specifies reasons for the rejected status.
type RejectedStatusReason21 struct {

	// Specifies the reason why the instruction or instruction cancellation has been rejected.
	ReasonCode *RejectedReason22Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason21) AddReasonCode() *RejectedReason22Choice {
	r.ReasonCode = new(RejectedReason22Choice)
	return r.ReasonCode
}

func (r *RejectedStatusReason21) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
