def slices(series, length):
    if length <= 0 or length > len(series):
        raise ValueError('slices("%s", %s)', (series, length))
    return [series[i : i + length] for i in range(len(series) - length + 1)]
