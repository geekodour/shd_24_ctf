default:
	@just --list

# Spin up any long running jobs for the local dev setup (eg.db)
# Once this is running we should be able to access it using: pgcli  -U postgres
# If we have corrent env vars
localsetup:
    nix run .#localsetup

buf:
	buf generate
