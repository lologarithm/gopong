require.config({
    baseUrl: '',
    paths: {
    	//REQUIRE LIBRARIES
        'jquery': '/js/bower_components/jquery/dist/jquery.min',
        'react': '/js/bower_components/react/react',
        'moment': '/js/bower_components/moment/min/moment.min',
        'history': '/js/bower_components/history.js/scripts/bundled/html5/native.history'
    },
});

require(['/js/components/views/index.js', '/js/utility/polyfills.js'], function (index, polyfills) {
	index();
});