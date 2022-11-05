function StartDrag(ev, targ) {
    const rect = ev.target.getBoundingClientRect();
  
    offsetX = ev.clientX - rect.x;
    offsetY = ev.clientY - rect.y;

    targ.parentNode.appendChild(targ);
};

function StopDrag(ev, targ) {
    document.getElementById(targ).style.left = ev.clientX - offsetX + 'px';
    document.getElementById(targ).style.top  = ev.clientY - offsetY + 'px';
};