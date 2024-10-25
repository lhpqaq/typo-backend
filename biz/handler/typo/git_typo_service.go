// Code generated by hertz generator.

package typo

import (
	"context"
	"fmt"

	"typo/biz/model/typo"
	"typo/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GitTypoMethod .
// @router /git-typo [POST]
func GitTypoMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req typo.GitTypoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(req)
	resp, err := service.CheckTypo(req.GitLink)
	if err != nil {
		fmt.Println(err)
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}
