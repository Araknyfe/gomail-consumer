// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "d2beb591327cb9dea1315e6d9fad90d2"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"4b39bc890dd6bd31e9602812cd40c17c": "1f8b08000000000000ff348e41cac2301046f7738a69d77f9a0bcc9f4d555c1415d185cbd14c69213192164442ee2eb1ba1adef0c17b542985e27974cd307b874a19a06ab56f4f97c31acbcb002d07480fc2d6005d837d19a0870144c4ad381730a56633c669deb1979c0b76fca33f7c8abb052f9f3d310e51faff3aa5e67cec72ae4d1beefd18fd12826c6d946922cd06487f2d407ad1968cd2f30e0000ffff3f5c8a31ba000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("templateBox", "./templates")
		b.SetResolver("welcome.html", packr.Pointer{ForwardBox: gk, ForwardPath: "4b39bc890dd6bd31e9602812cd40c17c"})
	}()
	return nil
}()
