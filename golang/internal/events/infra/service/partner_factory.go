package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerID int) (Partner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{
		partnerBaseURLs: partnerBaseURLs,
	}
}

func (partners *DefaultPartnerFactory) CreatePartner(partnerID int) (Partner, error) {
	baseURL, ok := partners.partnerBaseURLs[partnerID]
	if !ok {
		return nil, ErrPartnerNotFound(partnerID)
	}
	switch partnerID {
	case 1:
		return &Partner1{BaseURL: baseURL}, nil
	case 2:
		return &Partner2{BaseURL: baseURL}, nil
	default:
		return nil, ErrPartnerNotFound(partnerID)
	}
}

func ErrPartnerNotFound(partnerID int) error {
	return fmt.Errorf("partner with ID %d not found", partnerID)
}
