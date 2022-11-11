package zerandom

import "math/rand"

func Random(seed int) (int, error) {

	rand.Seed(int64(seed))
	i := rand.Int()

	return i, nil
}
