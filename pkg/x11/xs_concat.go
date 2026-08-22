package x11

func Xs_concat_words(words []string) XmString {
	var xmstr XmString
	if len(words) == 0 {
		xmstr = XmStringCreateLtoR("", XmSTRING_DEFAULT_CHARSET)
		return xmstr
	}

	for i, word := range words {
		if i > 0 {
			tmp := XmStringCreateLtoR(" ", XmSTRING_DEFAULT_CHARSET)
			xmstr = XmStringConcat(xmstr, tmp)
		}
		tmp := XmStringCreateLtoR(word, XmSTRING_DEFAULT_CHARSET)
		xmstr = XmStringConcat(xmstr, tmp)
	}
	return xmstr
}
