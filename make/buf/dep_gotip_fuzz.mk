# Must be set
$(call _assert_var,MAKEGO)
$(call _conditional_include,$(MAKEGO)/base.mk)
$(call _assert_var,CACHE_VERSIONS)
$(call _assert_var,CACHE_BIN)

# Settable
# https://github.com/golang/dl/commits/master 20211007 checked 20211028
GOTIP_VERSION ?= 6589945b0d1123571d5e8d78ca183133b535230f

GOTIP := $(CACHE_VERSIONS)/GOTIP/$(GOTIP_VERSION)
$(GOTIP):
	@rm -f $(CACHE_BIN)/gotip
	GOBIN=$(CACHE_BIN) go install golang.org/dl/gotip@$(GOTIP_VERSION)
	@rm -rf $(dir $@)
	@mkdir -p $(dir $@)
	@touch $@

# Settable
# https://go-review.googlesource.com/q/project:exp+branch:master+status:merged 20211028 checked 20211028
GOTIP_FUZZ_CL ?= 344955

GOTIP_FUZZ_HOME := $(CACHE)/gotip_fuzz
$(GOTIP_FUZZ_HOME):
	@mkdir -p $(GOTIP_FUZZ_HOME)

GOTIP_FUZZ := $(CACHE_VERSIONS)/GOTIP_FUZZ/$(GOTIP_FUZZ_CL)
$(GOTIP_FUZZ): $(GOTIP)
	(yes || true) | HOME=$(GOTIP_FUZZ_HOME) $(CACHE_BIN)/gotip download $(GOTIP_FUZZ_CL)
	@rm -rf $(dir $@)
	@mkdir -p $(dir $@)
	@touch $@
