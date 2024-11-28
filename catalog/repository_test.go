package catalog

import (
	"context"
	"reflect"
	"testing"

	elastic "gopkg.in/olivere/elastic.v5"
)

func TestNewElasticRepository(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewElasticRepository(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewElasticRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewElasticRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elasticRepository_Close(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			r.Close()
		})
	}
}

func Test_elasticRepository_PutProduct(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx context.Context
		p   Product
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			if err := r.PutProduct(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("elasticRepository.PutProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_elasticRepository_GetProductByID(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			got, err := r.GetProductByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("elasticRepository.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elasticRepository.GetProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elasticRepository_ListProducts(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx  context.Context
		skip uint64
		take uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			got, err := r.ListProducts(tt.args.ctx, tt.args.skip, tt.args.take)
			if (err != nil) != tt.wantErr {
				t.Errorf("elasticRepository.ListProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elasticRepository.ListProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elasticRepository_ListProductsWithIDs(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx context.Context
		ids []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			got, err := r.ListProductsWithIDs(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("elasticRepository.ListProductsWithIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elasticRepository.ListProductsWithIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_elasticRepository_SearchProducts(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx   context.Context
		query string
		skip  uint64
		take  uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &elasticRepository{
				client: tt.fields.client,
			}
			got, err := r.SearchProducts(tt.args.ctx, tt.args.query, tt.args.skip, tt.args.take)
			if (err != nil) != tt.wantErr {
				t.Errorf("elasticRepository.SearchProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("elasticRepository.SearchProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
