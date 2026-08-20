package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask23(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	c := CultureCampaign{DestinationCodes: []string{"TJ"}}
	require.NoError(t, s.CheckDestinationCoverage(context.Background(), c, []string{"TJ", "JS"}))
}
