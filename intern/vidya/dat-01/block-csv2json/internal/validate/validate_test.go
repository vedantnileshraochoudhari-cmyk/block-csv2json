package validate

import "testing"

func TestBlockNumber(t *testing.T) {
	_, err := BlockNumber("1001")
	if err != nil {
		t.Errorf("valid block number failed")
	}

	_, err = BlockNumber("abc")
	if err == nil {
		t.Errorf("invalid block number should fail")
	}
}

func TestTimestamp(t *testing.T) {
	_, err := Timestamp("1700000000")
	if err != nil {
		t.Errorf("valid timestamp failed")
	}

	_, err = Timestamp("invalid")
	if err == nil {
		t.Errorf("invalid timestamp should fail")
	}
}

func TestTxCount(t *testing.T) {
	_, err := TxCount("10")
	if err != nil {
		t.Errorf("valid tx_count failed")
	}

	_, err = TxCount("-5")
	if err == nil {
		t.Errorf("negative tx_count should fail")
	}
}
