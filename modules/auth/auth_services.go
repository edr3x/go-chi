package auth

import "github.com/edr3x/chi-explore/utils"

func testService(val string) (string, error) {
	if val == "err" {
		return "", utils.NewError(404, "Not found")
	}

	return val + " ok", nil
}
