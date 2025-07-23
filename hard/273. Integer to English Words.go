package hard

func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    a := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
    b := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
    c := []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
    d := []string{"", " Thousand", " Million", " Billion"}
    dot := -1
    ans := ""
    for num != 0 {
        dot += 1
        str := ""
        curr := num % 1000
        num = num / 1000
        if curr == 0 {
            continue
        }
        if curr % 100 / 10 == 1 {
            str = c[curr%10]
            curr = curr / 100
        } else {
            bit := curr % 10
            curr = curr / 10
            str = a[bit]
            bit = curr % 10
            curr = curr / 10
            if bit != 0 {
                str = b[bit] + " " + str
            }
        }
        if curr != 0 {
            str = a[curr] + " Hundred " + str
        }
        for i:= len(str)-1; i>=0 && str[i] == ' '; i-- {
            str = str[:i]
        }
        if ans == "" {
            ans = str + d[dot]
        } else {
            ans = str + d[dot] + " " + ans
        }
    }
    return ans
}
