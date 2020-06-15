type BrowserHistory struct {
    history []string
    current int
}


func Constructor(homepage string) BrowserHistory {
    h := BrowserHistory{}
    h.history = append(h.history, homepage)
    h.current = 0
    return h
}


func (this *BrowserHistory) Visit(url string)  {
    this.history = this.history[0:this.current+1]
    this.history = append(this.history, url)
    this.current++
}


func (this *BrowserHistory) Back(steps int) string {
    this.current -= steps
    if this.current < 0 {
        this.current = 0
    }
    return this.history[this.current]
}


func (this *BrowserHistory) Forward(steps int) string {
    this.current += steps
    if this.current >= len(this.history) {
        this.current = len(this.history)-1
    }
    return this.history[this.current]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
