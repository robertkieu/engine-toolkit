$(function(){

	// gentle scrolling
	// Anything with an href that points to something on the page
	// is gently scrolled to, rather than jumping.
	// Everything else is left alone.
	$("a[href]").click(function(e) { 
		var buffer = 20
		var dest = $(this).attr('href')
		dest = dest.substr(dest.indexOf('#')+1)
		var destEl = $('[id="'+dest+'"]')
		if (destEl.length > 0) {
			e.preventDefault()
			$('html,body').animate({ 
				scrollTop: destEl.offset().top - buffer
			}, 'slow')
			if (history.pushState) {
				history.pushState(null, null, '#'+dest);
			} else {
				location.hash = '#'+dest;
			}
		}
	})

})
