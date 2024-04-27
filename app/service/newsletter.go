package service

import (
	"context"
	"energy-response-assignment/app/repository"
	"energy-response-assignment/config"
	"energy-response-assignment/db/sqlc"
	"energy-response-assignment/util/mailer"
	"fmt"
	lop "github.com/samber/lo/parallel"
)

type Newsletter interface {
	SendNewsLetter(ctx context.Context, header string, body string) error
}

type newsletterImpl struct {
	cfg            *config.Config
	newsletterRepo repository.Newsletter
	mailDialer     mailer.Mailer
}

func (n *newsletterImpl) SendNewsLetter(ctx context.Context, header string, body string) error {
	limit := 100
	newsletter, err := n.newsletterRepo.CreateNewsLetter(ctx, sqlc.CreateNewsLetterParams{
		Header: header,
		Body:   body,
	})
	if err != nil {
		return err
	}
	activeSubscribersCount, err := n.newsletterRepo.CountActiveSubscriber(ctx)
	if err != nil {
		return err
	}
	for offset := 0; offset < int(activeSubscribersCount); offset += limit {
		subscribers, err := n.newsletterRepo.GetActiveSubscribersWithPaginate(ctx, sqlc.GetActiveSubscribersWithPaginateParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			return err
		}

		sendingEmailParams := lop.Map(subscribers, func(subscriber sqlc.Subscriber, _ int) sqlc.CreateSendingEmailsParams {
			return sqlc.CreateSendingEmailsParams{
				SubscriberID: subscriber.ID,
				NewsletterID: newsletter.ID,
			}
		})

		query := n.newsletterRepo.CreateSendingEmails(ctx, sendingEmailParams)
		var sendingEmails []sqlc.CreateSendingEmailsRow
		errs := make([]error, 0)
		query.Query(func(_ int, result []sqlc.CreateSendingEmailsRow, _ error) {
			sendingEmails = append(sendingEmails, result...)
			if err != nil {
				errs = append(errs, err)
			}
		})
		fmt.Println(sendingEmails, err)
		if len(errs) > 0 {
			errMsg := ""
			for _, err := range errs {
				errMsg += err.Error() + " | "
			}
			return fmt.Errorf(errMsg)
		}
		errChan := make(chan error)
		for _, sendingEmail := range sendingEmails {
			go func(errChan chan error, sendingEmail sqlc.CreateSendingEmailsRow) {
				status := sqlc.StatusCompleted
				err := n.mailDialer.SendMail([]string{sendingEmail.Email}, header, body)
				if err != nil {
					fmt.Println(err)
					status = sqlc.StatusFailed
				}
				_, err = n.newsletterRepo.UpdateSendingEmailsStatus(ctx, sqlc.UpdateSendingEmailsStatusParams{
					ID:     sendingEmail.ID,
					Status: status,
				})
				errChan <- err
			}(errChan, sendingEmail)
		}

		for range sendingEmails {
			err := <-errChan
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func NewNewsletter(cfg *config.Config, newsletterRepo repository.Newsletter, mailDialer mailer.Mailer) Newsletter {
	return &newsletterImpl{
		cfg:            cfg,
		newsletterRepo: newsletterRepo,
		mailDialer:     mailDialer,
	}
}
