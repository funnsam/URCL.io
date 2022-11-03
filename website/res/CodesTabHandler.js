function TabHandler(that, e) {
    if (e.key == 'Tab') {
        e.preventDefault();
        var start = that.selectionStart;
        var end = that.selectionEnd;
  
        that.value = that.value.substring(0, start) + "\t" + that.value.substring(end);
        that.selectionStart = that.selectionEnd = start + 1;
    }
};