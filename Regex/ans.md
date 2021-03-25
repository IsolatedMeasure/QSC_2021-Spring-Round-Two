

#### Q1：

```python
/[Aa]ndroid ((2\.3\.[1-7])|(2\.2\.[1-3])|(2\.0\.1)|(1[01])|(1\.[0156])|(2\.[0123])|(3\.[012])|(4\.((0(\.[1-4])?)|([1-4])))|([69]\.0)|(6\.0)|([578]\.[01])|(Pie)|(Oreo)|(Nougat)|(Marshmallow)|(Lollipop)|(KitKat)|(Jelly Bean)|(Ice Cream Sandwich)|(Honeycomb)|(Gingerbread)|(Froyo)|(Éclair)|(Donut)|(Cupcake)|[A-R])/
```



#### Q2：

```python
/^\* -?(([1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0)|(1(\.\d*)?[Ee]\d*)|([1-9](\d{0,2})(,\d\d\d)*(\.(\d\d\d,)*(\d{1,3}))?))$/
```

---

#### Q1：

```python
import re
pattern = re.compile(u"(\d\d\d)(.*)?\d\d\d(.*)?\d\d\d\d")
s = u'''415-555-1234	    415'''
print(pattern.search(s).group(1))

```

#### Q2:

```python
import re
pattern = re.compile(u"(?P<d>(((\w)|(\.))+))([-+.]\w+)*@(?P<b>(\w+))([-.]\w+)*\.\w+([-.]\w+)*")
s = u'''tom.riddle+regexone@hogwarts.com	        '''
r=pattern.search(s)
try :print(f"{r.group('d')}    {r.group('b')}")
except  : pass
```

#### Q3:

```python
import re
pattern = re.compile(
    u"(</(?P<a>([a-zA-Z]+))>)\s*$")
s = u'''<div class='test_style'><span>world</span>Test</div>	     '''
r = pattern.search(s)
try:
    print(f"{r.group('a')}")
except:
    pass
# (<((/([a-zA-Z]+))|(([a-zA-Z]+/)))>)\s*$
```

#### Q4：

```python
import re
pattern = re.compile(u"(?P<a>(.*))\.(?P<b>((mp4)|(jpg)|(jpeg)|(bmp)|(gif)|(tif)|(png))[^.])")
s = u'''updated_img0912.png	      updated_img0912       png        '''
r=pattern.search(s)
try : print(f"{r.group('a')}    {r.group('b')}") 
except  : pass
```

#### Q5：

```python
import re
pattern = re.compile(u"(?P<a>([a-zA-Z]+))://(?P<b>([a-zA-Z]+(.[a-zA-Z]+)*))(:(?P<c>([0-9]+)))?")
s = u'''ftp://file_server.com:21/top_secret/life_changing_plans.pdf  '''
r=pattern.search(s)
try : 
    print(f"{r.group('a')}    {r.group('b')}") 
    print(f"{r.group('c')}")
except  : 
    pass
```



#### Q6:

```python
import re
pattern = re.compile(
    u"widget\.List\.(?P<a>([a-zA-Z]+))\((?P<b>([a-zA-Z]+\.[a-zA-Z]+)):(?P<c>([0-9]+))\)")
s = u'''E/( 1553):   at widget.List.fillFrom(ListView.java:709) '''
r = pattern.search(s)
try:
    print(f"{r.group('a')}")
    print(f"{r.group('b')}")
    print(f"{r.group('c')}")
except:
    pass
```

