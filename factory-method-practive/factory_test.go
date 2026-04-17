package factory

import (
	"strings"
	"testing"
)

func TestTienMat(t *testing.T) {
	pt, err := LayPhuongThucThanhToan(TIEN_MAT)
	if err != nil {
		t.Fatal("Phải tạo được tiền mặt")
	}

	kq := pt.ThanhToan(10)
	if !strings.Contains(kq, "tiền mặt") {
		t.Error("Sai kết quả tiền mặt")
	}
}

func TestThe(t *testing.T) {
	pt, err := LayPhuongThucThanhToan(THE)
	if err != nil {
		t.Fatal("Phải tạo được thẻ")
	}

	kq := pt.ThanhToan(20)
	if !strings.Contains(kq, "thẻ") {
		t.Error("Sai kết quả thẻ")
	}
}

func TestSai(t *testing.T) {
	_, err := LayPhuongThucThanhToan(999)
	if err == nil {
		t.Error("Phải báo lỗi khi không tồn tại")
	}
}
