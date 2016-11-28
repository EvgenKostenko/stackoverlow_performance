### Сделать 2 API endpoint

1. **Топ 10 тегов которые встречаются в всех постах указанного автора**
    * Пользователь передается по Id
    * Должен иметь вид /toptags/{userID} где userID - Id юзера.
    * Теги должны быть возвращены в убывающем порядке по количеству постов с этим тегом    
    * Если есть 2 тега которые встречаются одинаково часто то стоит их расположить по алфавиту

-

2. **Топ 10 список авторов в чьих постах и комментариях встречаеться слово XXXX**
    * Должен иметь вид /topusers/{word} где word это слово
    * Слово должно встречаться целиком, минимальная длинна слова 3 символа
    * Если слово в одном посте или комментарии встречаться несколько раз это не повышает автора в топе, считаются только уникальные посты
    * Пользователи должны быть возвращены в убывающем порядке по количеству упоминания слова в постах и комментариях
    * Если у 2-х пользователей одинаково часто слово встречаться в постах и комментариях, то первым должен идти тот пользователь который зарегистрировался раньше


    
Формат ответа должен быть **json**
 
Endpoint должен возвращать массив из полных моделей данных, важно сохранить оригинальные названия полей и возвращать их все для того чтобы система тестирования приняла результат.

**Дамп базы данных**
https://drive.google.com/drive/folders/0B4edT1hr54MWTFc1WXEwY3J3ZVU?usp=sharing

**В данных есть пользователь с Id == -1 и посты с комментариями подписанные пользователем Id == 0 которого нет. Такие записи просьба игнорировать*

Примеры для тестов

ТОП тегов:

```
curl -X "GET" "http://localhost:8080/toptags/373"
```

```
[
  {
    "Count": 3391,
    "ExcerptPostId": 1970,
    "WikiPostId": 1969,
    "Id": 44,
    "TagName": "star-wars"
  },
  {
    "Count": 1515,
    "ExcerptPostId": 2381,
    "WikiPostId": 2380,
    "Id": 380,
    "TagName": "lord-of-the-rings"
  },
  {
    "Count": 166,
    "ExcerptPostId": 3089,
    "WikiPostId": 3088,
    "Id": 429,
    "TagName": "stargate-sg1"
  },
  {
    "Count": 130,
    "ExcerptPostId": 2224,
    "WikiPostId": 2223,
    "Id": 174,
    "TagName": "battlestar-galactica-2004"
  },
  {
    "Count": 116,
    "ExcerptPostId": 21609,
    "WikiPostId": 21608,
    "Id": 1138,
    "TagName": "darth-sidious"
  },
  {
    "Count": 116,
    "ExcerptPostId": 3993,
    "WikiPostId": 3992,
    "Id": 639,
    "TagName": "elves"
  },
  {
    "Count": 4,
    "ExcerptPostId": 27482,
    "WikiPostId": 27481,
    "Id": 1191,
    "TagName": "event-horizon"
  },
  {
    "Count": 29,
    "ExcerptPostId": 3075,
    "WikiPostId": 3074,
    "Id": 626,
    "TagName": "hyperspace"
  },
  {
    "Count": 23,
    "ExcerptPostId": 35282,
    "WikiPostId": 35281,
    "Id": 1474,
    "TagName": "istari"
  }
]
```

ТОП пользователей:

```
curl -X "GET" "http://localhost:8080/topusers/star"
```

```
[
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 20774,
    "Age": 0,
    "DisplayName": "Valorum",
    "AboutMe": "",
    "Reputation": 192109,
    "LastAccessDate": "2016-06-12T02:06:34.913",
    "DownVotes": 3661,
    "AccountId": 3776439,
    "Location": "UK",
    "Views": 19302,
    "CreationDate": "2013-12-26T15:54:24.750",
    "ProfileImageUrl": "https://i.stack.imgur.com/ZopOw.jpg"
  },
  {
    "WebsiteUrl": "http://stackoverflow.com/",
    "UpVotes": 0,
    "Id": 976,
    "Age": 44,
    "DisplayName": "DVK-in-exile",
    "AboutMe": "<p><strong>Areas of expertise:</strong></p>\n\n<ul>\n<li>Perl expert (specifically enterprise software development\n<li>Sybase (table design, query optimization and stored proc development)\n<li>Java (including Reflection and Javascript embedding)\n</ul>\n\n<p><strong>Areas of familiarity:</strong> </p>\n\n<ul>\n<li>Web programming (EmbPerl, JSP, CSS, HTML, JavaScript)\n<li>C++\n<li>Shell scripting, *nix system administration\n</ul>\n",
    "Reputation": 208393,
    "LastAccessDate": "2016-06-12T02:44:18.867",
    "DownVotes": 2167,
    "AccountId": 41067,
    "Location": "New York, NY",
    "Views": 14321,
    "CreationDate": "2011-02-28T04:00:56.213",
    "ProfileImageUrl": "https://www.gravatar.com/avatar/a7efd5c795e9797c30fce4126daaa9c0?s=128&d=identicon&r=PG&f=1"
  },
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 40294,
    "Age": 0,
    "DisplayName": "Praxis",
    "AboutMe": "<p><em>\"There's a science fiction in the space between you and me.\"</em></p>\n\n<p>Tracy Chapman, \"Telling Stories\"</p>\n\n<hr>\n\n<h1>Me</h1>\n\n<ul>\n<li>Physicist</li>\n<li>Born and mainly raised in the UK, living in Canada these days</li>\n<li>Ardent supporter of the Oxford comma</li>\n</ul>\n\n<hr>\n\n<h1>Me and Sci-Fi</h1>\n\n<p>I like:</p>\n\n<ul>\n<li><em>Star Trek</em></li>\n<li><em>Star Wars</em></li>\n<li><em>The X-Files</em></li>\n<li><em>The Matrix</em></li>\n<li><em>Babylon 5</em></li>\n<li><em>Buffyverse</em></li>\n<li><em>Doctor Who</em></li>\n<li><em>The Terminator</em> and <em>Terminator 2</em></li>\n<li><em>Nowhere Man</em> (does anyone else like this?)</li>\n<li><em>Twin Peaks</em> (anything by Lynch, actually)</li>\n<li>David Cronenberg's films</li>\n<li>Lots of other stuff</li>\n</ul>\n\n<p>I dislike:</p>\n\n<ul>\n<li>The JJ Abrams <em>Star Trek</em> films (<strike>although I am quite hopeful about <em>The Force Awakens</em></strike>)</li>\n<li>The <em>Star Wars</em> prequel trilogy and George Lucas' \"touch-ups\" to the original trilogy</li>\n<li>New <em>Terminator</em> films (they should have quit while they were ahead)</li>\n</ul>\n\n<hr>\n\n<h1>Me and <em>Star Trek</em></h1>\n\n<p>I've watched every episode of <em>Star Trek</em> (including the animated series) and have watched every film.  I grew up watching <em>TNG</em> when it was still on the air.  I have also read plenty of Star Trek novels and comics.  I have a penchant for non-standard Starfleet insignia (e.g. insignia appearing in alternate realities or timelines).</p>\n\n<p>The best gift I ever received was the <em>Star Trek: The Next Generation Technical Manual</em> by Okuda and Sternbach (the year it was published).</p>\n\n<p>I feel that <em>Star Trek</em> played a large role in inspiring me toward a career in science, even if a lot of <em>Trek</em> science is just technobabble.</p>\n\n<hr>\n\n<h1>Me and You</h1>\n\n<p>I have a lot of questions and maybe even some answers to give.  I hope to learn new things and to be a productive member of this community.</p>\n",
    "Reputation": 60223,
    "LastAccessDate": "2016-06-12T03:17:23.887",
    "DownVotes": 358,
    "AccountId": 5596321,
    "Location": "Canada",
    "Views": 2851,
    "CreationDate": "2015-01-08T17:31:21.503",
    "ProfileImageUrl": "https://i.stack.imgur.com/onKUP.jpg?s=128&g=1"
  },
  {
    "WebsiteUrl": "http://hubcityblues.com",
    "UpVotes": 0,
    "Id": 2765,
    "Age": 0,
    "DisplayName": "Thaddeus Howze",
    "AboutMe": "<p>Thaddeus Howze is a California-based technologist and author who has worked with computer technology since the 1980’s doing graphic design, computer science, programming, network administration and IT leadership.</p>\n\n<p>His non-fiction work has appeared in numerous magazines: Huffington Post, Gizmodo, Black Enterprise, the Good Men Project, Examiner.com, The Enemy, Panel &amp; Frame, Science X, Loud Journal, ComicsBeat.com, and Astronaut.com. He maintains a diverse collection of non-fiction at his blog, A Matter of Scale.</p>\n\n<p>Thaddeus is a popular and well-read writer on the Q&amp;A site Quora.com in over fifty various subjects. He is also a moderator and contributor to the Science Fiction and Fantasy Stack Exchange with over fourteen hundred articles in a four year period.</p>\n\n<p>He is an author and contributor at Scifiideas.com. His speculative fiction has appeared online at Medium.com, ScifiIdeas.com, and the Au Courant Press Journal. He has a wide collection of his work on his website, Hub City Blues. His recently published works can be found here. He also maintains a wide collection of his writing and editing work on Medium.</p>\n\n<p><a href=\"http://i.stack.imgur.com/NCt49.png\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/NCt49.png\" alt=\"enter image description here\"></a></p>\n\n<p><strong><a href=\"http://hubcityblues.com/about/thaddeus-howze/\" rel=\"nofollow\">Speculative Fiction Author</a>, <a href=\"http://linkedin.com/in/ebonstorm\" rel=\"nofollow\">IT Professional</a>, <a href=\"https://ebonstorm.contently.com\" rel=\"nofollow\">Journalist</a></strong></p>\n",
    "Reputation": 157823,
    "LastAccessDate": "2016-06-12T02:20:12.243",
    "DownVotes": 80,
    "AccountId": 893673,
    "Location": "Hayward, CA",
    "Views": 5746,
    "CreationDate": "2011-09-06T19:47:37.910",
    "ProfileImageUrl": "https://i.stack.imgur.com/KycFn.jpg?s=128&g=1"
  },
  {
    "WebsiteUrl": "http://scifi.stackexchange.com/users/22917/n-soong",
    "UpVotes": 0,
    "Id": 22917,
    "Age": 20,
    "DisplayName": "Often Right",
    "AboutMe": "<h1>Selected Achievements on this site</h1>\n\n<ul>\n<li><a href=\"http://meta.scifi.stackexchange.com/questions/7131/honorary-rank-for-this-site\">Science Fiction Fantasy Rank</a>: Admiral - that means you can\ncall me 'Admiral @Often Right' if you want!</li>\n<li>Third user to get the <a href=\"http://scifi.stackexchange.com/help/badges/139/star-trek\">gold [tag:star-trek] badge</a></li>\n<li>Second user to get the <a href=\"http://scifi.stackexchange.com/help/badges/131/star-trek-tng\">silver [tag:star-trek-tng] badge</a></li>\n<li>Second user to get the <a href=\"http://scifi.stackexchange.com/help/badges/177/star-trek-data\">bronze [tag:star-trek-data] badge</a></li>\n<li>I've asked <a href=\"http://scifi.stackexchange.com/search?tab=votes&amp;q=user%3A22917%20%5Bstar-trek%5D%20is%3Aquestion\">the most questions for the [tag:star-trek] tag</a></li>\n<li>In the top 20 users in terms of reputation</li>\n<li>Earned bronze tag badges for <a href=\"/questions/tagged/star-trek-ds9\" class=\"post-tag\" title=\"show questions tagged &#39;star-trek-ds9&#39;\" rel=\"tag\">star-trek-ds9</a>, <a href=\"/questions/tagged/star-trek-tos\" class=\"post-tag\" title=\"show questions tagged &#39;star-trek-tos&#39;\" rel=\"tag\">star-trek-tos</a> and <a href=\"/questions/tagged/star-trek-voy\" class=\"post-tag\" title=\"show questions tagged &#39;star-trek-voy&#39;\" rel=\"tag\">star-trek-voy</a></li>\n<li>95% answer acceptance rate on my questions</li>\n</ul>\n",
    "Reputation": 39172,
    "LastAccessDate": "2016-06-11T05:51:02.073",
    "DownVotes": 20,
    "AccountId": 3864077,
    "Location": "Omicron Theta",
    "Views": 2606,
    "CreationDate": "2014-02-16T22:25:07.170",
    "ProfileImageUrl": "https://i.stack.imgur.com/qDYpR.jpg?s=128&g=1"
  },
  {
    "WebsiteUrl": "http://None",
    "UpVotes": 0,
    "Id": 44025,
    "Age": 38,
    "DisplayName": "Wad Cheber",
    "AboutMe": "<p><em>The eggplant will salute other pigeons!</em></p>\n\n<p>SF&amp;F's self-styled expert on <em>John Carpenter's The Thing</em> (1982).</p>\n\n<p><a href=\"http://i.stack.imgur.com/vAzyI.gif\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/vAzyI.gif\" alt=\"enter image description here\"></a></p>\n\n<p>The guy who asked about every single detail related to Tolkien's work. <em>The Thing</em>, the original <em>Star Wars</em> trilogy, <em>Rick and Morty</em>, <em>Terminator</em>, <em>Aliens</em>, Stephen King's <em>Dark Tower</em> series, the book version of <em>I Am Legend</em>, <em>The Road</em>, <em>The Walking Dead</em>, Max Brooks' books <em>World War Z</em> and <em>The Zombie Survival Guide</em>, and other zombie fare are all fair game too.  Zombie fanatic.  I <strong>very</strong> rarely read fiction, but feel free to peruse my Goodreads shelves.</p>\n\n<p><a href=\"https://www.goodreads.com/review/list/7136149?page=1&amp;shelf=currently-reading\" rel=\"nofollow\">https://www.goodreads.com/review/list/7136149?page=1&amp;shelf=currently-reading</a></p>\n\n<p><a href=\"http://i.stack.imgur.com/hyoFIm.gif\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/hyoFIm.gif\" alt=\"enter image description here\"></a></p>\n\n<p><em>Mayday, mayday.  Can anyone hear me?  Over.  This is U.S. Station 31.  Do you read me?  We found something in the ice.  We need some help down here.  Can anybody hear me?  We found something ... we found something ... we found something ...</em><br>\n- <a href=\"http://www.youtube.com/watch?v=nVXkTNGNSh4\" rel=\"nofollow\">Outpost 31</a></p>\n\n<p>\"Act in such a way that you treat humanity, whether in your own person or in the person of any other, never merely as a means to an end, but always at the same time as an end.”<br>\n- Immanuel Kant, <em>Grounding for the Metaphysics of Morality</em> </p>\n\n<p>\"Life's a piece of shit, when you look at it;  life's a laugh and death's a joke its true.  You'll see its all a show, keep 'em laughing as you go; just remember that the last laugh is on you, and  always look on the bright side of life...\"<br>\n- Monty Python</p>\n\n<p>\"Talking about music is like painting about farting.\"<br>\n- Pickles the Drummer</p>\n\n<p>\"It's all a big nothing.  What makes you think you're so special?\"<br>\n- Livia Soprano</p>\n",
    "Reputation": 28220,
    "LastAccessDate": "2016-06-12T03:32:13.477",
    "DownVotes": 829,
    "AccountId": 6095099,
    "Location": "USA",
    "Views": 3471,
    "CreationDate": "2015-04-06T20:48:19.910",
    "ProfileImageUrl": "https://i.stack.imgur.com/VveKd.jpg?s=128&g=1"
  },
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 931,
    "Age": 0,
    "DisplayName": "Evil Angel",
    "AboutMe": "<p><a href=\"http://i.stack.imgur.com/abL1U.gif\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/abL1U.gif\" alt=\"Emma Stone with Lipstick\"></a></p>\n\n<p><a href=\"http://i.stack.imgur.com/9SuL4.jpg\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/9SuL4.jpg\" alt=\"Lips Lips Lips\"></a></p>\n",
    "Reputation": 1,
    "LastAccessDate": "2016-01-20T06:40:12.810",
    "DownVotes": 522,
    "AccountId": 281536,
    "Location": "",
    "Views": 15183,
    "CreationDate": "2011-02-23T03:05:34.000",
    "ProfileImageUrl": "https://i.stack.imgur.com/RGsyr.jpg?s=128&g=1"
  },
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 22250,
    "Age": 0,
    "DisplayName": "Hypnosifl",
    "AboutMe": "",
    "Reputation": 31564,
    "LastAccessDate": "2016-06-12T02:04:27.737",
    "DownVotes": 70,
    "AccountId": 3938459,
    "Location": "",
    "Views": 740,
    "CreationDate": "2014-01-30T15:56:38.730",
    "ProfileImageUrl": "https://www.gravatar.com/avatar/f901fa15361d7c7c5282c33878f4bc96?s=128&d=identicon&r=PG&f=1"
  },
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 5184,
    "Age": 0,
    "DisplayName": "phantom42",
    "AboutMe": "<p><a href=\"http://i.stack.imgur.com/Kb0yO.jpg\" rel=\"nofollow\"><img src=\"http://i.stack.imgur.com/Kb0yO.jpg\" alt=\"it&#39;s over now... the music of the night...\"></a></p>\n",
    "Reputation": 82069,
    "LastAccessDate": "2016-06-11T22:04:39.017",
    "DownVotes": 835,
    "AccountId": 1170648,
    "Location": "ZZ9 Plural Z Alpha",
    "Views": 4219,
    "CreationDate": "2012-03-06T21:45:06.980",
    "ProfileImageUrl": "https://i.stack.imgur.com/OQ9iv.png?s=128&g=1"
  },
  {
    "WebsiteUrl": "",
    "UpVotes": 0,
    "Id": 2242,
    "Age": 28,
    "DisplayName": "Izkata",
    "AboutMe": "<p>Completely unintentional, but as of Aug 20, 2012:  <code>this user has accepted an answer for 7 of 9 eligible questions</code>.</p>\n",
    "Reputation": 44900,
    "LastAccessDate": "2016-06-12T03:30:09.043",
    "DownVotes": 255,
    "AccountId": 234199,
    "Location": "Chicago, IL",
    "Views": 1992,
    "CreationDate": "2011-06-18T18:58:26.810",
    "ProfileImageUrl": ""
  }
]
```
