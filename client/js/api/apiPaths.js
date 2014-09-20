define([], function () {
	$.ajaxSetup({
		success: function (resp) {
			console.log(resp);
		}
	});

	return {
		ACCOUNT: {
			GET_ACCOUNT: function () { return '/api/account'},
			UPDATE_ACCOUNT: function (region) { return '/api/account?region=' + region},
		},
		SUMMONER : {
			GET_MATCH_HISTORY : function(name, region){return '/api/summoner/matchHistory?name=' + name + '&region=' + region},
			GET_RANKED_DATA : function(name, region){return '/api/summoner/rankedData?name=' + name + '&region=' + region},
			GET_LFSCORES : function(summonerId, region){return '/api/summoner/LFScores?summonerId=' + summonerId + '&region=' + region}
		},
		CHAMPION : {

		},
		MATCH : {
			GET_MATCH : function(matchId, summonerId, region){ return '/api/match?matchId=' + matchId + '&summonerId=' + summonerId + '&region=' + region }
		},
		TOURNAMENTS : {
			GET_TOURNAMENTS : function () { return '/api/summoner/tournaments' },
			GET_TOURNAMENT : function (id) { return id ? '/api/tournament?id=' + id : '/api/tournament'; },
			GENERATE_BRACKETS : function (id) { return id ? '/api/tournament/generateBrackets?id=' + id : '/api/tournament/generateBrackets'; },
			ADVANCE_TOURNAMENT : function (id) { return id ? '/api/tournament/advance?id=' + id : '/api/tournament/advance'; }
		}
	};
})