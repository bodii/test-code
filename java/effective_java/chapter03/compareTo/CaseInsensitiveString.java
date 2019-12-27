package compareTo;

import java.lang.Comparable;

public final class CaseInsensitiveString implements Comparable<CaseInsensitiveString> {
    public String s = "";
    public int compareTo(CaseInsensitiveString cis) {
        return String.CASE_INSENSITIVE_ORDER.compare(s, cis.s);
    }
}