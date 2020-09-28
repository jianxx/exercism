def saddle_points(matrix):
    if any(len(matrix[0]) != len(r) for r in matrix):
        raise ValueError("irregular matrix")
    points = []
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            if matrix[i][j] == max(matrix[i]) and matrix[i][j] == min(
                [row[j] for row in matrix]
            ):
                points.append({"row": i + 1, "column": j + 1})
    return points
