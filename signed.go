package randomorg

func (r *Random) GenerateSignedIntegers(n int, min, max int64) (
	[]int64,
	map[string]interface{},
	string,
	error,
) {
	if n < 1 || n > 1e4 {
		return nil, nil, "", ErrParamRange
	}
	if min < -1e9 || min > 1e9 || max < -1e9 || max > 1e9 {
		return nil, nil, "", ErrParamRange
	}

	params := map[string]interface{}{
		"n":   n,
		"min": min,
		"max": max,
	}

	values, random, sig, err := r.requestSignedCommand("generateSignedIntegers", params)
	if err != nil {
		return nil, nil, "", err
	}

	ints := make([]int64, len(values))
	for i, value := range values {
		f := value.(float64)
		ints[i] = int64(f)
	}

	return ints, random, sig, nil
}
