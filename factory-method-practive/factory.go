package factory

import (
	"errors"
	"fmt"
)

type PhuongThucThanhToan interface {
	ThanhToan(soTien float32) string
}

type TienMat struct{}
type TheNganHang struct{}

const (
	TIEN_MAT = 1
	THE      = 2
)

func LayPhuongThucThanhToan(loai int) (PhuongThucThanhToan, error) {
	switch loai {
	case TIEN_MAT:
		return new(TienMat), nil
	case THE:
		return new(TheNganHang), nil
	default:
		return nil, errors.New(fmt.Sprintf("Không hỗ trợ phương thức %d", loai))
	}
}

func (t TienMat) ThanhToan(soTien float32) string {
	return fmt.Sprintf("%.2f thanh toán bằng tiền mặt", soTien)
}

func (t TheNganHang) ThanhToan(soTien float32) string {
	return fmt.Sprintf("%.2f thanh toán bằng thẻ", soTien)
}
