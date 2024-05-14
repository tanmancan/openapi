GEN_CMD:=/usr/local/bin/docker-entrypoint.sh

SPEC_FILES:=$(wildcard ./spec/*/*.yaml)
SPEC_DIR:=$(dir $(SPEC_FILES))

TMP_DIR:=$(dir tmp/)
OUT_DIR:=$(dir $(TMP_DIR)openapi/)

GIT_HOST:=github.com
GIT_REMOTE_URL:=$(GIT_REMOTE_URL)
GIT_SSH_HOST:=git@github.com
GIT_PATH:=$(patsubst $(GIT_SSH_HOST):%.git,%,$(GIT_REMOTE_URL))
GIT_USER:=$(word 1, $(subst /, ,$(GIT_PATH)))
GIT_REPO_ID:=$(word 2, $(subst /, ,$(GIT_PATH)))

all: clean $(OUT_DIR) 

$(SPEC_FILES):
	@echo "generating OpenAPI client from $(@)"

	$(GEN_CMD) generate \
		--git-host $(GIT_HOST) \
		--git-user-id $(GIT_USER) \
		--git-repo-id $(GIT_REPO_ID)/$(patsubst spec/%/,%,$(dir $@)) \
		-i $@ \
		-g go \
		-o $(OUT_DIR)$(patsubst spec/%/,%,$(dir $@)) \
		--additional-properties=packageName=$(patsubst spec/%/,%,$(dir $@)) \

	@rm -rf $(patsubst spec/%/,%,$(dir $@))
	@cp -r $(OUT_DIR)$(patsubst spec/%/,%,$(dir $@)) .
	@rm -rf $(TMP_DIR)

$(OUT_DIR): 
	@mkdir -p $(TMP_DIR)
	@make $(SPEC_FILES)

clean:
	rm -rf $(OUT_DIR)

gohelp:
	$(GEN_CMD) config-help -g go

.PHONY: $(SPEC_FILES)