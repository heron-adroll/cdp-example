package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/heron-adroll/cdp-example/graph/generated"
	"github.com/heron-adroll/cdp-example/graph/model"
)

func (r *accountResolver) CompanyAttributes(ctx context.Context, obj *model.Account) (*model.CompanyAttributes, error) {
	// We already fetched company attributes for the domains in accountsByAttribute but since this is a resolver and not a regular field
	// we can't assign the value. We need to fetch all the attributes from the Account again
	return &model.CompanyAttributes{
		Industry: "industry1",
	}, nil
}

func (r *accountResolver) SurgeRecords(ctx context.Context, obj *model.Account) (*model.IntentSurgeRecords, error) {
	return &model.IntentSurgeRecords{
		Topic: "topic1",
	}, nil
}

func (r *advertisableResolver) Accounts(ctx context.Context, obj *model.Advertisable, talEid *string, tagEid *string) (*model.TalAccountSource, error) {
	return &model.TalAccountSource{
		Accounts: []*model.Account{
			{Domain: "domain1"},
			{Domain: "domain2"},
		},
	}, nil
}

func (r *attributesAccountSourceResolver) AccountsByIntent(ctx context.Context, obj *model.AttributesAccountSource, intentFilters *model.IntentFilters) (*model.IntentAccountSource, error) {
	var accountsByIntent []*model.Account
	for _, acc := range obj.Accounts {
		if acc.Domain == "domain4" {
			accountsByIntent = append(accountsByIntent, acc)
		}
	}

	return &model.IntentAccountSource{
		Accounts: accountsByIntent,
	}, nil
}

func (r *intentAccountSourceResolver) AccountsByAttributes(ctx context.Context, obj *model.IntentAccountSource, datFilters *model.DatFilters) (*model.AttributesAccountSource, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Advertisable(ctx context.Context, eid string) (*model.Advertisable, error) {
	return &model.Advertisable{Eid: "eid"}, nil
}

func (r *talAccountSourceResolver) AccountsByAttributes(ctx context.Context, obj *model.TalAccountSource, datFilters *model.DatFilters) (*model.AttributesAccountSource, error) {
	// talEID := "get from obj context"
	// tagEID := "get from obj context"

	// accounts = slargmaAPI.accounts_by_dat(talEID, tagEID, datFilters)

	// return &model.AttributesAccountSource{
	// 	Accounts: []*model.Account{
	// 		[Domain: accounts[0].domain,
	// 		[Domain: accounts[1].domain],
	// 	}
	// }, nil
	// check comments on line 15
	return &model.AttributesAccountSource{
		Accounts: []*model.Account{
			{Domain: "domain3"},
			{Domain: "domain4"},
		},
	}, nil
}

func (r *talAccountSourceResolver) AccountsByIntent(ctx context.Context, obj *model.TalAccountSource, intentFilters *model.IntentFilters) (*model.IntentAccountSource, error) {
	panic(fmt.Errorf("not implemented"))
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Advertisable returns generated.AdvertisableResolver implementation.
func (r *Resolver) Advertisable() generated.AdvertisableResolver { return &advertisableResolver{r} }

// AttributesAccountSource returns generated.AttributesAccountSourceResolver implementation.
func (r *Resolver) AttributesAccountSource() generated.AttributesAccountSourceResolver {
	return &attributesAccountSourceResolver{r}
}

// IntentAccountSource returns generated.IntentAccountSourceResolver implementation.
func (r *Resolver) IntentAccountSource() generated.IntentAccountSourceResolver {
	return &intentAccountSourceResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TalAccountSource returns generated.TalAccountSourceResolver implementation.
func (r *Resolver) TalAccountSource() generated.TalAccountSourceResolver {
	return &talAccountSourceResolver{r}
}

type accountResolver struct{ *Resolver }
type advertisableResolver struct{ *Resolver }
type attributesAccountSourceResolver struct{ *Resolver }
type intentAccountSourceResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type talAccountSourceResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *advertisableResolver) AccountsByTal(ctx context.Context, obj *model.Advertisable, talEid *string, tagEid *string) (*model.TalAccountSource, error) {
	panic(fmt.Errorf("not implemented"))
}
