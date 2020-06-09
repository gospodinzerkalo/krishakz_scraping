<h1>Krisha KZ scraping</h1>
<h2>Review</h2>
<p>Web scraping website krisha.kz. In Progress....</p><hr>
<h2>Using</h2>
<p><b style="color: #08ff20">[GET]</b>  /allRent - get all apartments for rent</p><hr>
<p><b style="color: #08ff20">[GET]</b>  /allSell - get all apartments for sale</p><hr>
<p><b style="color: #08ff20">[GET]</b>  /sell/{params} - get apartments by parameters</p>
<h5>Parameters:</h5>
<p> <b>room</b> - number of rooms. Values = [1,2,3,4,5.100]. Ex: /sell/room=2</p>
<p> <b>price_from</b> - starting price. Values: min=1; max=15000000000000. Ex: /sell/price_from=100000000</p>
<p> <b>price_to</b> - the final price. Values: min=1; max=15000000000000. Ex: /sell/price_to=100000000</p>
<p> <b>has_photo</b> - select apartment only which has photo. Values: 1. Ex: /sell/has_photo=1</p>
<p> <b>checked</b> - select only verified apartments. Values: 1. Ex: /sell/checked=1</p>
<p> <b>owner</b> - select apartments only by owner (not realtor). Values: 1. Ex: /sell/owner=1</p>
<p> <b>building</b> - select apartments by type of building. Values: 1 - brick, 2 - panel, 3 - monolithic,4 - another. Ex: /sell/building=3</p>
<p> <b>floor_from</b> - starting floor of apartment. Values: min=0; max=500. Ex: /sell/floor_from=3</p>
<p> <b>floor_to</b> - the end floor of apartment. Values: min=0; max=500. Ex: /sell/floor_to=5</p>
<p> <b>year_from</b> - start year of construction. Values: min=1850; max=2030. Ex: /sell/year_from=2000</p>
<p> <b>year_to</b> - the end year of construction. Values: min=1850; max=2030. Ex: /sell/year_from=2015</p>
<p> <b>toilet</b> - type of bathroom. Values: 1 - separated; 2 - combined; 3 - more than 2; 4 - not. Ex: /sell/toilet=1</p>
<p> <b>priv_dorm</b> - former dormitory. Values: 1 - yes; 2 - no. Ex: /sell/priv_dorm=2</p>
<h6>Note</h6>
<p>You can use all the parameters you want via &. Example: /sell/room=2&checked=1&price_from=8000000&price_to=15000000</p>

<hr>
<h2>Dependencies</h2>
<p><a href="github.com/valyala/fasthttp">fasthttp</a></p>
<p><a href="github.com/buaazp/fasthttprouter">fasthttprouter</a></p>
<p><a href="github.com/PuerkitoBio/goquery">goquery</a></p>
<p><a href="github.com/urfave/cli">cli</a></p> <hr>

<h2>Clone the project</h2>
<code>git clone https://github.com/gospodinzerkalo/krishakz_scraping</code> <hr>
<h2>Install dependencies</h2>
<code>make depends</code>

<h2> Build and Run </h2>
<code>make build</code><br>
<code>make run</code>

