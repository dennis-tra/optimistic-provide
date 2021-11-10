def fmt_thousands(val: int) -> str:
    return format(val, ",")


def fmt_percentage(total: int):
    return lambda val: "%.1f%%" % (100 * val / total)
