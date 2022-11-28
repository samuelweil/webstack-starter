sa_acct=--iam-account=auto-tester@web-stack-starter.iam.gserviceaccount.com
key_file=test-serviceaccount-key

run:
	go run cmd/api.go

$(key_file):
	gcloud iam service-accounts keys create $(key_file) $(sa_acct)

clean:
	gcloud iam service-accounts keys delete $(shell jq '.private_key_id' $(key_file)) $(sa_acct)
	rm $(key_file)

integration-tests: $(key_file)
	GOOGLE_APPLICATION_CREDENTIALS=$(key_file) go test ./test -count=1