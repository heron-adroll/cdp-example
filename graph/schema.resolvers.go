package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/heron-adroll/cdp-example/graph/generated"
	"github.com/heron-adroll/cdp-example/graph/model"
)

func (r *accountResolver) CompanyAttributes(ctx context.Context, obj *model.Account) (*model.CompanyAttributes, error) {
	// check context for map[obj.Domain][companyAttributes][firstParty]|[thirdParty]
	return &model.CompanyAttributes{
		FirstParty: model.FirstPartyCompanyAttributes{Industry: "industry1"},
		ThirdParty: model.ThirdPartyCompanyAttributes{Employees: 10},
	}, nil
}

func (r *accountResolver) SurgeRecords(ctx context.Context, obj *model.Account) (*model.IntentSurgeRecords, error) {
	// check context for map[obj.Domain][surgeRecords]
	return &model.IntentSurgeRecords{
		Topic: "topic1",
	}, nil
}

func (r *accountResolver) References(ctx context.Context, obj *model.Account) (*model.AccountReferences, error) {
	// call references service
	campaign := "campaign1"
	return &model.AccountReferences{
		Campaigns: []*string{&campaign},
	}, nil
}

func (r *advertisableResolver) Accounts(ctx context.Context, obj *model.Advertisable, talEid *string, tagEid *string) (*model.TalAccountSource, error) {
	// call slargma and fetch TAL accounts
	return &model.TalAccountSource{
		Accounts: []*model.Account{
			{Domain: "domain1"},
			{Domain: "domain2"},
		},
	}, nil
}

func (r *attributesAccountSourceResolver) AccountsByIntent(ctx context.Context, obj *model.AttributesAccountSource, domains *model.DomainsFilter, searchType *string, intentFilters *model.IntentFilters) (*model.IntentAccountSource, error) {
	var accountsByIntent []*model.Account
	for _, acc := range obj.Accounts {
		if acc.Domain == "domain2" {
			accountsByIntent = append(accountsByIntent, acc)
		}
	}

	return &model.IntentAccountSource{
		Accounts: accountsByIntent,
	}, nil
}

func (r *intentAccountSourceResolver) AccountsByAttributes(ctx context.Context, obj *model.IntentAccountSource, domains *model.DomainsFilter, searchType *string, firstPartyFilters *model.FirstPartyFilters, thirdPartyFilters *model.ThirdPartyFilters) (*model.AttributesAccountSource, error) {
	accounts := []*model.Account{}
	for _, acc := range obj.Accounts {
		if acc.Domain == "domain1" {
			accounts = append(accounts, acc)
		}
	}

	return &model.AttributesAccountSource{
		Accounts: accounts,
	}, nil
}

func (r *queryResolver) Advertisable(ctx context.Context, eid string) (*model.Advertisable, error) {
	return &model.Advertisable{Eid: "eid"}, nil
}

func (r *talAccountSourceResolver) AccountsByAttributes(ctx context.Context, obj *model.TalAccountSource, domains *model.DomainsFilter, searchType *string, firstPartyFilters *model.FirstPartyFilters, thirdPartyFilters *model.ThirdPartyFilters) (*model.AttributesAccountSource, error) {
	accounts := obj.Accounts
	return &model.AttributesAccountSource{
		Accounts: accounts,
	}, nil
}

func (r *talAccountSourceResolver) AccountsByIntent(ctx context.Context, obj *model.TalAccountSource, domains *model.DomainsFilter, searchType *string, intentFilters *model.IntentFilters) (*model.IntentAccountSource, error) {
	return &model.IntentAccountSource{
		Accounts: obj.Accounts,
	}, nil
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
type companyAttributesResolver struct{ *Resolver }
