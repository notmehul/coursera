<h3>
Concurrency in Go
</h3>
<h2 style="text-align: left;">
Week 1 -&nbsp;Why Use Concurrency?</h2>
<h3 style="text-align: left;">
Parallel Execution</h3>
<br>
<ul style="text-align: left;">
<li>big property of Go language is that concurrency is built in the language</li>
<li>concurrency and parallelism are two closely related ideas</li>
<li>in C, Python etc...you can do concurrent programming but you need to use some 3rd party library&nbsp;</li>
</ul>
<h4 style="text-align: left;">
Parallel Execution</h4>
<ul style="text-align: left;">
<li>parallel is not the same as concurrent</li>
<li>why do we need concurrency?</li>
<li>parallel execution is when two programs are executed at the same time</li>
<li>at some point in time instructions from two programs are executing in parallel</li>
<ul>
<li>at time t and instruction is being performed for both P1 and P2</li>
</ul>
<li>one processor core is executing 1 instruction at a time</li>
<li>if we want to have parallel execution, we need two processors or at least two processor cores</li>
<li>we need replicated hardware: e.g. CPU1 and CPU2</li>
<li>or if we have quad-core CPU then we can run 4 instructions in parallel, at the same time</li>
</ul>
<br>
<h4 style="text-align: left;">
Why use Parallel Execution?</h4>
<br>
<ul style="text-align: left;">
<li>tasks may be completed more quickly</li>
<li>you get a better throughput overall</li>
<li>example: two piles of dishes to wash</li>
<ul>
<li>two dishwashers can complete twice as fast as one</li>
</ul>
<li>BUT: some tasks must be performed sequentially</li>
<ul>
<li>example: wash dish then dry dish; it has to be in this order</li>
<ul>
<li>must wash before you can dry</li>
</ul>
</ul>
<li>some tasks are more parallelizable then other</li>
<ul>
<li>we can't parallelize washing and drying of the same dish - even if we have hundred dishwashers and sinks available</li>
</ul>
<li>some things can't be parallelized</li>
</ul>
<h4 style="text-align: left;">
Von Neumann Bottleneck</h4>
<div>
<ul style="text-align: left;">
<li>running concurrent product is hard</li>
<li>students usually learn sequential programming at least at undergraduate curriculum</li>
<li>concurrent programming is difficult</li>
<li>can we achieve speedup without Parallelism?</li>
<li>Solution 1: Design faster processors</li>
<ul>
<li>get speedup without changing software</li>
<li>this is what used to be case till recently&nbsp;</li>
<li>hardware has been getting faster and faster but this really stopped now</li>
<li>faster = processor clock rate; this used to get doubled every year or so</li>
</ul>
<li>limitation on a speed that we have now is called Von Neumann Bottleneck: delayed access to memory</li>
<ul>
<li>CPU has to access memory to get instructions and data</li>
<li>memory is always slower than CPU</li>
<li>even if CPU has high clock rate, it has to wait on memory access; lots of time is wasted just waiting for memory access</li>
</ul>
<li>Solution 2:&nbsp; design processors with more memory</li>
<ul>
<li>build cache, fast memory on chip</li>
<li>that's what's traditionally been done till now</li>
<li>again, the same software, the same code would run faster</li>
</ul>
<li>But this is not the case anymore</li>
</ul>
<h4 style="text-align: left;">
Moore's Law</h4>
</div>
<div>
<ul style="text-align: left;">
<li>it is not a physical law but more an observation of trends which used to be actual in past</li>
<li>Predicted that transistor density would double every two years</li>
<li>smaller transistors switch faster</li>
<li>exponential increase in density would lead to exponential increase in software speed</li>
<li>this is not the case anymore so software engineers had to do something</li>
</ul>
<h4 style="text-align: left;">
Power Wall</h4>
<h4 style="text-align: left;">
Power/Temperature Problem</h4>
<div>
<ul style="text-align: left;">
<li>the speedup that we get from Moore's Law can't continue because transistors consume power and power is becoming a critical issue - Power Wall</li>
<li>transistors consume power whenever they switch&nbsp;</li>
<li>as you increase number of transistors, power used also increases</li>
<ul>
<li>smaller transistors use less power but transistor density scaling is much faster</li>
</ul>
<li>nowadays many devices are portable and are using batteries; it is a limited power available</li>
<li>even if devices is plugged to a socket, the issue is a temperature</li>
<ul>
<li>we need fans and heat sink above processors e.g. i7; with out them it would melt</li>
<li>high power leads to high temperature</li>
<li>air cooling (fans) can only remove so much heat</li>
</ul>
</ul>
<h4 style="text-align: left;">
Dynamic Power</h4>
</div>
<div>
<ul style="text-align: left;">
<li>generic equation for Power: <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">P = alpha * C * F * V^2</span></li>
<ul>
<li>alpha - percent of time switching; transistors consume dynamic power when they switch (go from 0 to 1 or from 1 to 0; when they don't switch they don't use dynamic power)</li>
<li>C - capacitance (related to size; goes down as transistor shrinks)</li>
<li>F - clock frequency (we want to increase frequency)</li>
<li>V - voltage swing (from low to high) (0 is 0V and 1 is 5V; we want to reduce the voltage in order to reduce power e.g. 1 to be 0.3V)</li>
</ul>
</ul>
</div>
<h4 style="text-align: left;">
Dennard Scaling</h4>
<div>
<ul style="text-align: left;">
<li>it is paired together with Moore's Law</li>
<li>Voltage should scale with transistor size</li>
<ul>
<li>smaller transistor =&gt; voltage should be smaller</li>
</ul>
<li>We want to scale down voltage to keep power consumption law</li>
<li>We want to have Dennard Scaling</li>
<li>but it can't go forever:&nbsp;</li>
<li>Problem: Voltage can't go too low</li>
<ul>
<li>must stay above threshold voltage otherwise transistor won't switch</li>
<li>noise problems occur; if there is a noise in a system, voltage can go + or - a volt or two and we can get wrong readings/errors</li>
</ul>
<li>Problem: doesn't consider leakage power</li>
<ul>
<li>leakage happens when you have thin insulators</li>
<li>as you scale everything down, insulators get thinner and leakage increases</li>
</ul>
<li>Dennard scaling must stop</li>
</ul>
<h4 style="text-align: left;">
Multi-Core Systems</h4>
</div>
<div>
<ul style="text-align: left;">
<li>generic equation for Power:&nbsp;<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">P = alpha * C * F * V^2</span></li>
<ul>
<li><span style="font-family: inherit;">we can't increase frequency</span></li>
<li>frequency still goes up but with very slow rate</li>
</ul>
<li>we then increase cores on chips, without increasing frequency =&gt; multi-core systems</li>
<ul>
<li>they are still increasing density</li>
</ul>
<li>if you have 4 core system but are running program only on one core, you are wasting 3 other cores who are idling</li>
<li>parallel execution is needed to exploit multi-core systems</li>
<li>code made to execute on multiple cores</li>
<ul>
<li>we want to divide our program into smaller chunks which are executed in parallel, on different cores</li>
</ul>
<li>different programs on different cores</li>
<li>parallel compilers:&nbsp;</li>
<ul>
<li>they take sequential code and parallelize it&nbsp;</li>
<li>they chop it in tasks that can be run in parallel</li>
<li>this is extremely complex problem</li>
<li>this does not work that well</li>
</ul>
<li>concurrent programming: it is actually a programmer who divides code into tasks that will be run in parallel</li>
</ul>
<h3 style="text-align: left;">
Concurrent vs Parallel</h3>
</div>
<h4 style="text-align: left;">
Concurrent Execution</h4>
<div>
<ul style="text-align: left;">
<li>concurrent execution is not necessarily the same as parallel execution</li>
<li><b>concurrent</b>: start and end times overlap</li>
<ul>
<li>they don't literally execute at the same time</li>
<li>if we take any point in time we can see that either task1 or task2 is executing</li>
<li>concurrent: tasks are competing for chunks of time (CPU clocks) which would execute their own instruction, not of the other competitor task</li>
<li>task1: --------&nbsp; &nbsp; &nbsp; &nbsp;-------&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;(time --&gt;)</li>
<li>task2:&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;-----&nbsp;</li>
</ul>
<li><b>parallel</b>: execute at exactly the same time</li>
<ul>
<li>if we take any point in time we can see that both task1 and task2 are executing</li>
<li>task1: ----------------&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;(time --&gt;)</li>
<li>task2:&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;-----&nbsp;</li>
</ul>
<li>completion time for both tasks is longer in case of concurrent case</li>
<li>why do concurrent execution? why just not do task2 after task1 if joint time of execution is the same?</li>
</ul>
<h4 style="text-align: left;">
Concurrent vs Parallel</h4>
<ul style="text-align: left;">
<li>parallel tasks need to be executed on different hardware (different cores)</li>
<li>concurrent tasks may be executed on the same hardware (e.g. on one core)</li>
<ul>
<li>only one task is actually executed at a time</li>
</ul>
<li>Mapping from tasks to hardware (which task is executed on which core) is not directly controlled by the programmer</li>
<ul>
<li>at least not in Go</li>
</ul>
</ul>
<h4 style="text-align: left;">
Concurrent Programming</h4>
</div>
<div>
<ul style="text-align: left;">
<li>which tasks are executed on which core</li>
<li>programmer determines which tasks can be executed in parallel; e.g. task1 has to do some task before before task2 etc...</li>
<li>programmer describes what can be (not what will be) executed in parallel</li>
<ul>
<li>programmer defines possible concurrency</li>
</ul>
<li>what will be executed in parallel depends on how tasks are mapped to hardware</li>
<li>mapping task to hardware is done by:</li>
<ul>
<li>operating system</li>
<li>Go runtime scheduler</li>
</ul>
</ul>
<h4 style="text-align: left;">
Hiding Latency</h4>
<ul style="text-align: left;">
<li>let's say we do concurrent programming and have one core; so why bother doing concurrency at all?</li>
<li>even though we can't do parallel execution we can still get significant performance improvements because tasks typically have to periodically wait for some slow events (e.g. input/output)</li>
<li>when CPU's have to communicate with memory, network, file system, video card...this is all slow</li>
<li>example: reading from memory</li>
<ul>
<li>X = Y + Z</li>
<li>CPU has to read Y and Z from memory, do operation and write result back to memory</li>
<li>to read/write from/to memory CPU has to wait hundreds of cycles but in the meantime CPU could use those clocks to execute some useful tasks</li>
</ul>
<li>Other concurrent tasks can operate while one task is waiting</li>
</ul>
</div>
<h4 style="text-align: left;">
Hardware Mapping&nbsp;</h4>
<div>
<ul style="text-align: left;">
<li>Parallel Execution:</li>
<ul>
<li>task1 --&gt; Core1</li>
<li>task2 --&gt; Core2</li>
</ul>
<li>Concurrent Execution:</li>
<ul>
<li>task1 ---/--&gt; Core1</li>
<li>task2 --/</li>
</ul>
</ul>
<h4 style="text-align: left;">
Hardware Mapping in Go</h4>
<ul style="text-align: left;">
<li>not under direct control of programmer in Go or any other standard language</li>
<li>programmer makes parallelism possible</li>
<li>this is very hard task which would slow down programming</li>
<li>hard task for it depends on many factors</li>
<ul>
<li>e.g. underlying hardware architecture - programmers should not care about details about it</li>
</ul>
<li>simple arbitrary multi-core system:</li>
</ul>
<div>
core&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; core&nbsp;</div>
<div>
&nbsp; &nbsp;|&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; |</div>
cache ------------- cache</div>
<div>
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; |</div>
<div>
&nbsp; &nbsp; &nbsp; &nbsp;shared memory</div>
<div>
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; |</div>
<div>
cache ------------- cache</div>
<div>
&nbsp; &nbsp;|&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; |</div>
</div>
<div>
core&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; core&nbsp;</div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>cache - a local cache for each core</li>
<li>one big consideration when figuring out hardware mapping is where is the data?</li>
<ul>
<li>if core 1 has to perform some task we'd like data to be in its cache&nbsp;</li>
<li>if that data is in cache of core 2 then core 2 should perform that task</li>
<ul>
<li>task has to be performed on that core where data is</li>
</ul>
</ul>
<li>in Go we define which task can be executed in parallel</li>
</ul>
</div>
<div>
<br></div>
<br>
<h2 style="text-align: left;">
Week 2 -&nbsp;CONCURRENCY BASICS</h2>
<h3 style="text-align: left;">
Processes</h3>
<div>
<ul style="text-align: left;">
<li>a lot of concurrent execution ideas came from operating systems</li>
</ul>
<h4 style="text-align: left;">
Processes</h4>
</div>
<div>
<ul style="text-align: left;">
<li>an instance of a running program</li>
<li>things unique to a process</li>
<ul>
<li>Memory</li>
<ul>
<li>virtual address space&nbsp;</li>
<li>code&nbsp;</li>
<li>stack - region of memory which usually handles function calls</li>
<li>heap - for memory allocations</li>
<li>shared libraries - shared between processes</li>
</ul>
<li>Registers - they store 1 value; tiny, super-fast, memory</li>
<ul>
<li>Program counter - tells what instruction is executed now or the next one</li>
<li>data registers</li>
<li>stack pointer</li>
<li>...</li>
</ul>
</ul>
</ul>
<h4 style="text-align: left;">
Operating System</h4>
</div>
<div>
<ul style="text-align: left;">
<li>Allows many processes to execute concurrently</li>
<ul>
<li>makes sure that virtual address space is not overlapping</li>
<li>makes sure that all processes get fair share of processor time and resources</li>
<li>these processes are run concurrently; they switch quickly...like 20ms</li>
<li>from user's perspective it looks all processes run in parallel although they don't - OS ensures the illusion of parallelism</li>
</ul>
</ul>
<h4 style="text-align: left;">
Task Manager</h4>
</div>
<div>
<ul style="text-align: left;">
<li>shows all running processes</li>
<ul>
<li>foreground</li>
<li>background</li>
</ul>
</ul>
</div>
<h4 style="text-align: left;">
Scheduling</h4>
<div>
<ul style="text-align: left;">
<li>one of the main tasks of OS</li>
<li>OS schedules processes for execution</li>
<li>Gives the illusion of parallel execution</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">...</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">process1</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">process2</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">process3</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">process1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">process2</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">...</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>OS gives fair access to CPU, memory, etc...</li>
<li>there are many different scheduling algorithms</li>
<ul>
<li>one is called Round Robin - processes simply alternate in a round fashion (1, 2, 3, 1, 2, 3, 1, 2, 3...) so every process gets the same chunk of time</li>
<li>if processes don't have the same priority - processes with high priority would get more CPU time, they would be scheduled more frequently</li>
</ul>
<li>embedded systems: some tasks are critical e.g. breaking - that would be considered a high-priority task while playing stereo music would be a low-priority task</li>
</ul>
</div>
<h4 style="text-align: left;">
Context Switch</h4>
<div>
<ul style="text-align: left;">
<li>control flow changes from one process to another</li>
<li>e.g. switching from processA to processB</li>
<li>before each switch OS has to save the state of currently running process and restore it when next time its execution gets resumed</li>
<li>this state is called the <b>context </b>- all the stuff unique to the process we listed before</li>
<li>process "context" must be swapped</li>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">...</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">processA</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">context switch</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">processB</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">context switch</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">processA</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">...</span></div>
<div>
<ul style="text-align: left;">
<li>during context switch periods kernel of the OS is running</li>
<li>context switch usually happens after a timer times out</li>
</ul>
<h4 style="text-align: left;">
Threads vs Processes</h4>
</div>
<div>
<ul style="text-align: left;">
<li>there used to be only Processes</li>
<li>downsize of processes: process switching time is long (switching between processes is slow because memory access)</li>
<li>to speed up this: threads</li>
<li>thread is like a process but ig has less contexts; it shares some of the context with other threads in a process</li>
<li>Parts of Process context shared among process threads:</li>
<ul>
<li>Virtual memory</li>
<li>File descriptors</li>
</ul>
<li>Specific (unique) parts of Process context for each thread:</li>
<ul>
<li>Stack</li>
<li>Data registers</li>
<li>Code (PC)</li>
</ul>
<li>Switching between threads is faster because there is less context - less data that has to be read/written from/to memory</li>
</ul>
</div>
<h4 style="text-align: left;">
Goroutines</h4>
<div>
<ul style="text-align: left;">
<li>goroutine is basically a like a thread in Go</li>
<li>many goroutines execute within a single OS thread</li>
<li>Go takes a process with main thread and schedules / switches goroutines within that thread</li>
</ul>
<h4 style="text-align: left;">
Go Runtime Scheduler</h4>
</div>
<div>
<ul style="text-align: left;">
<li>schedules goroutines inside an OS thread (main thread)</li>
<li>like a little OS inside a single OS thread</li>
<li>Go runs on a main thread and switches goroutines on one thread</li>
<li>from OS point of view - there is only a single thread</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; Main thread</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;|</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp;Logical processor</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;|</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; Go runtime</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;|</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;|--------------------------&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;/&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; |&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;\</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Goroutine1&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;Goroutine2&nbsp; &nbsp; &nbsp; &nbsp; Goroutine3</span></div>
<div>
<br></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>Go runtime scheduler uses <b>Logical Processor</b> which is mapped to a thread</li>
<li>typically there is one Logical Processor which is mapped to a main thread</li>
<li>since all these goroutines are running on one thread, we don't have parallelism but concurrency</li>
<li>we can increase number of Logical Processors - mapped to different threads and OS can map those threads to different cores</li>
<li>program can determine how many Logical Processors will be there; default is 1 (so we'll have concurrent execution of routines) but can be increased (so we might have parallel goroutines execution - if OS schedules running different threads on different cores)</li>
</ul>
<h4 style="text-align: left;">
Interleavings</h4>
</div>
<div>
<ul style="text-align: left;">
<li>writing concurrent code is hard as it's difficult mentally to keep track what's happening on which thread</li>
<li>the overall state of the machine is not deterministic</li>
<li>in case of crash, it can happen at different places</li>
<li>order of execution within task is known</li>
<li>order of execution between concurrent tasks is unknown</li>
<li>Let's look instructions in two tasks:</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1&nbsp;</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">1: a = b+ c&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">2: d = e + f</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">3: g = h + i</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">1: r = s + t</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">2: u = v + w</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">3: x = y + z</span></div>
<div>
<br></div>
<h4 style="text-align: left;">
Possible Interleavings</h4>
<div>
<br></div>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">1: a = b+ c&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;1: r = s + t</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">2: d = e + f</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;2: u = v + w</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">3: g = h + i</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;3: x = y + z</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">OR</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">1: a = b+ c&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">2: d = e + f</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">3: g = h + i</span></div>
</div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;1: r = s + t</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;2: u = v + w</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;3: x = y + z</span></div>
</div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>many interleavings are possible</li>
<li>must consider all possibilities</li>
<li>ordering is non-deterministic</li>
</ul>
</div>
<h4 style="text-align: left;">
Race Conditions</h4>
<div>
<ul style="text-align: left;">
<li>problem that can happen because of these interleavings that can happen</li>
<li>the outcome of the program depends on the interleaving; interleavings are indeterministic =&gt; outcome is undeterministic</li>
<li>the outcome of the program depends on non-deterministic ordering</li>
<li>interleaving can change every time we run a program</li>
<li>we want to have determinism: for the same set of inputs we want the same set of outputs</li>
<li>we want outcome of the program does not depend on interleavings</li>
</ul>
<div>
<span style="font-family: inherit;">1st running - 1st interleaving combination</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;print x</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = x + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: inherit;">Output: 1</span></div>
<div>
<span style="font-family: inherit;"><br></span></div>
<div>
<span style="font-family: inherit;">2nd running - 2nd interleaving combination</span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = x + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;print x</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: inherit;">Output: 2</span></div>
</div>
<div>
<ul style="text-align: left;">
<li>This needs to be avoided, prevented</li>
<li>Races occur due to communications: two tasks are communicating through the share variable x</li>
<li>if we didn't have this communication there would not be race condition</li>
<li>communication between tasks is often unavoidable&nbsp;</li>
</ul>
<h4 style="text-align: left;">
Communication Between Tasks</h4>
</div>
<div>
<ul style="text-align: left;">
<li>Threads are largely independent but not completely independent</li>
<li>Web server, one thread per client</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Web&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;Client 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">page&nbsp; --&gt; Web server --&gt;&nbsp; Client 2</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Data&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Client 3</span></div>
<div>
<ul style="text-align: left;">
<li>Clients are coming at the same time</li>
<li>Example: webpage shows visits counter</li>
<li>Image processing example: 1 thread per pixel block</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">thread1 --&gt; [][] --&gt; thread2</span></div>
</div>
<div>
<ul style="text-align: left;">
<li><span style="font-family: inherit;">image processing is parallelizable but there must be some level of communication between threads e.g. in case of blurring</span></li>
</ul>
<div>
<br></div>
</div>
<h2 style="text-align: left;">
Week 3 -&nbsp;THREADS IN GO</h2>
<h3 style="text-align: left;">
Goroutines</h3>
<div>
<ul style="text-align: left;">
<li>to create threads of execution we have to use some of constructs built in Go</li>
</ul>
</div>
<h4 style="text-align: left;">
Creating a goroutine</h4>
<div>
<ul style="text-align: left;">
<li>one goroutine is created automatically to execute the main()</li>
<li>other goroutines are created using the <b><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">go</span></b> keyword</li>
<li>code snippet from some <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">main()</span> function with only one goroutine (the one for <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">main()</span>)</li>
<ul>
<li>a is assigned 2 only after <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">foo()</span> returns</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">foo() </span>blocks <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">main()</span></li>
</ul>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">foo()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a = 2&nbsp;</span></div>
<div>
<ul style="text-align: left;">
<li>code snippet from some <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">main()</span> function with two goroutines</li>
<ul>
<li>with <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">go foo()</span> we create a new goroutine</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a</span> might be (but also might not be)&nbsp;&nbsp;assigned before <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">foo()</span> returns - this depends on how the scheduler would schedule concurrent goroutine</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">foo()</span> is non-blocking for <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">main()</span><span style="font-family: inherit;">- it continues execution without waiting for foo() to return</span></li>
</ul>
</ul>
</div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><b><span style="color: blue;">go</span></b> foo()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a = 2&nbsp;</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<h4 style="text-align: left;">
<span style="font-family: inherit;">Exiting a goroutine</span></h4>
<div>
<ul style="text-align: left;">
<li><span style="font-family: inherit;">goroutine exits when code of the associated with its function is complete</span></li>
<li><span style="font-family: inherit;">when the main goroutine is complete, all other goroutines exit, even if they are not finished</span></li>
<li><span style="font-family: inherit;">a goroutine may not complete its execution because main completes early</span></li>
</ul>
<h3 style="text-align: left;">
Exiting goroutines</h3>
</div>
<div>
<ul style="text-align: left;">
<li>goroutines are forced to exit when main goroutine exits</li>
</ul>
<h4 style="text-align: left;">
Early Exit</h4>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go fmt.Printf("New routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Printf("Main routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>we don't know the order of the execution of these routines - this is undeterministic</li>
<li>we'd expect to see both messages but actually only the 2nd message is printed (almost always)</li>
<li>this is because the scheduler seems to give preference to main routine (this might not be completely true) and also (very true) main() does not block after first message is scheduled to be printed in new goroutine</li>
<li>main() is finished before the new goroutine has a chance to start</li>
<li>we want main() to wait for other goroutines to complete</li>
</ul>
<h4 style="text-align: left;">
Delayed Exit</h4>
</div>
<div>
<ul style="text-align: left;">
<li>the following snippet shows a hacky/bad solution which <i>might </i>work:</li>
</ul>
</div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go fmt.Printf("New routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;<b><span style="color: blue;">time.Sleep</span></b>(100 * time.Millisecond)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Printf("Main routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
</div>
<div>
<ul style="text-align: left;">
<li>we put main goroutine to sleep so another goroutine can resume or start&nbsp;</li>
<li>now is printed 1st and then 2nd message</li>
<li>this is a hack, bad solution because we assumed that 100ms would be enough for 2nd goroutine to be completed/executed. But this might be not the case, we don't know how much time it would take. Maybe it would take 100ms for OS to schedule main Go application thread to some other thread and we assume that this would not happen which is bad! Maybe the Go runtime schedules another goroutine - and we assumed it won't - again, bad!</li>
<li>timing is non-deterministic</li>
<li>we need formal&nbsp;<b>synchronization&nbsp;constructs</b></li>
</ul>
<h3 style="text-align: left;">
Basic Synchronization</h3>
</div>
<h4>
Synchronization</h4>
<div>
<ul style="text-align: left;">
<li>synchronization is when multiple threads agree on a timing of an event</li>
<li>there are global events whose execution is viewed by all threads, simultaneously</li>
<li>one goroutine does not know the timing of other goroutines</li>
<li>synchronization breaks that - it introduces some global events that every thread sees at the same time</li>
<li>this is important in order to restrict interleavings</li>
<li>e.g. two possible interleavings showing a rave condition: output depends on the interleaving; interleavings/schedules are non-deterministic</li>
</ul>
<div>
<br></div>
</div>
<div>
<div>
<span style="font-family: inherit;">1st running - 1st interleaving combination</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1&nbsp; &nbsp; Task 2&nbsp;&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">------&nbsp; &nbsp; ------</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;print x</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = x + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: inherit;">Output: 1</span></div>
<div>
<span style="font-family: inherit;"><br></span></div>
<div>
<span style="font-family: inherit;">2nd running - 2nd interleaving combination</span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = x + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;print x</span></div>
</div>
</div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>we want to know what is the intention of the programmer</li>
<li>let's assume they want printing to happen after x in incremented</li>
<li>synchronization is used to restrict bad interleavings</li>
<li>synchronization:</li>
<ul>
<li>we need some global event (GLOBAL EVENT) so both tasks/threads/gorotuines can see at the same time</li>
<li>Task 1: event happens</li>
<li>Task 2: we need a conditional execution:</li>
<ul>
<li>if that event happened, goroutine will wait or run;</li>
<li>if GLOBAL EVENT happened, we'll execute printing</li>
</ul>
</ul>
</ul>
</div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; Task 2&nbsp;&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">------&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; ------</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x = x + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><b><span style="color: blue;">GLOBAL EVENT</span></b></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;if <b><span style="color: blue;">GLOBAL EVENT</span></b></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; print x</span></div>
</div>
<div>
<ul style="text-align: left;">
<li>synchronization is the opposite of concurrency</li>
<li>synchronization makes some threads/goroutines to wait</li>
<li>synchronization is reducing effective use of hardware</li>
<li>but in some cases it is necessary - when things have to happen in order</li>
<li>synchronization is a necessary evil</li>
</ul>
</div>
<h3 style="text-align: left;">
Wait Groups</h3>
<div>
<ul style="text-align: left;">
<li>type of synchronization that are common</li>
</ul>
</div>
<h4 style="text-align: left;">
Sync WaitGroup</h4>
<div>
<ul style="text-align: left;">
<li>sync package contains functions to synchronize between goroutines</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">sync.WaitGroup</span> forces a goroutine to wait for other goroutines</li>
<li>wait group is like a group of goroutines that our goroutine has to wait for</li>
<ul>
<li>our goroutine will not continue until all goroutines from WaitGroup finish</li>
</ul>
<li>We can wait on 1 or more other goroutines</li>
<li>in this example: we want to make main goroutine to wait for 2nd goroutine</li>
</ul>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go fmt.Printf("New routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Printf("Main routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
</div>
</div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>WaitGroup contains an internal counter</li>
<ul>
<li>increment a counter for each goroutine we want to wait for</li>
<ul>
<li>if there are 3 gorountines, we'll increase it by 3</li>
</ul>
<li>decrement a counter when each goroutine completes</li>
<li>waiting goroutine has to wait till this counter becomes 0</li>
</ul>
</ul>
<h4 style="text-align: left;">
Using WaitGroup</h4>
</div>
<div>
<ul style="text-align: left;">
<li>main() runs in main goroutine</li>
<li>foo() runs in 2nd (worker) goroutine</li>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Main thread:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;var wg sync.WaitGroup</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.<b><span style="color: blue;">Add</span></b>(1)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go foo(&amp;wg)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.<b><span style="color: blue;">Wait</span></b>()</span></div>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Foo thread:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;</span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">wg.</span><b style="font-family: &quot;Courier New&quot;, Courier, monospace;"><span style="color: blue;">Done</span></b><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">()</span></div>
<div>
<br></div>
</div>
<div>
<ul style="text-align: left;">
<li>WaitGroup methods:</li>
<ul>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Add()</span> increments the counter</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Done()</span> decrements the counter</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Wait()</span> blocks until counter == 0</li>
</ul>
</ul>
</div>
<h4 style="text-align: left;">
<span style="font-family: inherit;">WaitGroup Example</span></h4>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func foo(wg *sync.WaitGroup) {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Printf("New routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.<b><span style="color: blue;">Done</span></b>()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;var wg sync.WaitGroup</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.<b><span style="color: blue;">Add</span></b>(1)</span></div>
<div>
&nbsp; &nbsp; &nbsp; &nbsp;<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">go foo(&amp;wg)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.<b><span style="color: blue;">Wait</span></b>()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;</span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">fmt.Printf("Main routine")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">)</span></div>
<div>
<ul style="text-align: left;">
<li><span style="font-family: inherit;">Output:&nbsp;"New routine" and then "Main routine"</span></li>
</ul>
<br><ul style="text-align: left;">
</ul>
</div>
<h3 style="text-align: left;">
Communication</h3>
<h4 style="text-align: left;">
<br></h4>
<h4 style="text-align: left;">
Goroutine Communication</h4>
<ul style="text-align: left;">
<li>goroutine can wait for each other</li>
<li>but they can also communicate with each other</li>
<li>generally, goroutines work together to perform a bigger task</li>
<li>these goroutines are not completely independent</li>
<li>they are doing a small piece of a bigger task</li>
<li>e.g. web server</li>
<ul>
<li>makes sense to create a new thread for each new browser connection</li>
<li>all these threads share some data; e.g. if some data is sent from one browser, can be seen by another browser =&gt; these threads have to cooperate</li>
</ul>
<li>example: find the product of 4 integers</li>
<ul>
<li>make 2 goroutines, each multiplies a pair</li>
<li>main goroutine multiplies the 2 results</li>
<li>need to send ints from main goroutine to two go subroutines</li>
<li>need to send results from subroutines back to main routine</li>
</ul>
</ul>
<h4 style="text-align: left;">
Channels</h4>
<div>
<ul style="text-align: left;">
<li>used for communication between goroutines</li>
<li>used to transfer data between goroutines</li>
<li>channels are typed</li>
<ul>
<li>one channel can handle integers, another strings etc...</li>
</ul>
<li>use <b><span style="color: blue; font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">make()</span></b> to create a channel:</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c := make(chan int)</span></div>
</div>
<ul style="text-align: left;">
<li>send and receive data using <b><span style="color: blue;">arrow operator (<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&lt;-</span>)</span></b></li>
<ul>
<li>send data on a channel:&nbsp;</li>
</ul>
</ul>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c &lt;- 3</span><br>
<ul style="text-align: left;"><ul>
<li>receive data from a channel:&nbsp;</li>
</ul>
</ul>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x := &lt;- c</span><br>
<h4 style="text-align: left;">
Channel Example</h4>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">func prod(v1 int, v2 int, c chan int) {</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;c &lt;- v1 * v2</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">}</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;"><br></span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">func main() {</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;c := make(chain int)</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;go prod(1, 2, c)</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;go prod(3, 4, c)</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;a := &lt;- c</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;b := &lt;- c</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">&nbsp; &nbsp;fmt.Println(a * b)</span></div>
<div style="text-align: left;">
<span style="font-family: Courier New, Courier, monospace;">}</span></div>
<div style="text-align: left;">
<br></div>
<div style="text-align: left;">
</div>
<ul style="text-align: left;">
<li><span style="font-family: Courier New, Courier, monospace;">a</span> gets whatever first comes out of the channel</li>
<li><span style="font-family: Courier New, Courier, monospace;">b</span> gets whatever second comes out of the channel</li>
<li>There is also another way to send data between goroutines - via passing arguments to it when it is starting</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h3 style="text-align: left;">
Blocking in Channels</h3>
<h4 style="text-align: left;">
Unbuffered Channel</h4>
<div>
<ul style="text-align: left;">
<li>by default, when channel is created (with <b><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">make()</span></b>), it is unbuffered</li>
<ul>
<li>default is unbuffered</li>
</ul>
<li>unbuffered channels can't hold data in transit</li>
<li>the implications are:&nbsp;</li>
<ul>
<li>sending blocks until the data is received</li>
<li>receiving blocks until data is sent</li>
</ul>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1: c &lt;- 3</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">one hour later...</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 2: x := &lt;- c</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<h4 style="text-align: left;">
</h4>
<h4 style="text-align: left;">
Blocking and Synchronization</h4>
<ul style="text-align: left;">
<li>this is also doing a synchronization, just like a WaitGroup</li>
<ul>
<li>Task2 has to wait till Task1 sends data</li>
</ul>
<li>channel communication is synchronous</li>
<li>Blocking is the same as waiting for communication</li>
<li>this kind of communication can be used for pure synchronization - we can freely drop the data - throw away received result:</li>
</ul>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1: c &lt;- 3</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">one hour later...</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 2: &lt;- c</span></div>
</div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>Task 2 is receiving the data but is throwing it away</li>
<li>All we do here is synchronizing two tasks: Task 2 has to wait for Task 1 to send the data</li>
<li>this is another way to implement WaitGroup's Wait()</li>
</ul>
<div>
<br></div>
</div>
<h4 style="text-align: left;">
Buffered Channel&nbsp;</h4>
<h4 style="text-align: left;">
Channel Capacity</h4>
<div>
<ul style="text-align: left;">
<li>channels by default are unbuffered</li>
<ul>
<li>they have no capacity to hold the data</li>
<li>unbuffered channels have capacity 0</li>
</ul>
<li>channels can have some capacity</li>
<li>capacity is the number of objects channel can hold in transit</li>
<li>to make channel a buffered, we can use an optional argument of make() function - it's second argument would define a channel capacity</li>
</ul>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c := make(chan int, 3)</span></div>
<div>
<ul style="text-align: left;">
<li>default size is 0</li>
<li>channel with some capacity still blocks under some conditions</li>
<li><u>sending </u>only blocks if <u>buffer is full</u></li>
<ul>
<li>e.g. if we have 3 sends with no receives, new sends will be blocked</li>
<li>as soon as a new receive happens, the next, 4th send will unblock</li>
</ul>
<li><u>receiving</u> only blocks if <u>buffer is empty</u> - if there is nothing in a buffer, channel read operation will block until there is something to be read from the channel</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Channel Blocking, Receive</h4>
</div>
<div>
<ul style="text-align: left;">
<li>channel with capacity 1</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1 -------&gt; [&nbsp; &nbsp; ] --------&gt; Task 2</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c &lt;- 3</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a := &lt;- c</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">b := &lt;- c&nbsp;</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>first receive blocks (in Task 2) until send occurs (in Task 1)</li>
<li>second receive blocks forever (in Task 2)</li>
</ul>
<br><ul style="text-align: left;">
</ul>
</div>
<div>
<h4>
Channel Blocking, Send</h4>
</div>
<div>
<div>
<ul>
<li>channel with capacity 1</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1 -------&gt; [&nbsp; &nbsp; ] --------&gt; Task 2</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task 1:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c &lt;- 3</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">c &lt;- 4</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">a := &lt;- c</span></div>
<div>
<ul>
<li>second send blocks till first receive is done</li>
</ul>
<div>
<br></div>
</div>
</div>
<h4 style="text-align: left;">
Use of Buffering</h4>
<div>
<ul style="text-align: left;">
<li>buffering is used when producer and consumer work in different speeds</li>
<li>sender and receiver do not need to operate at exactly the same speed</li>
<li>Producer is generating data&nbsp;</li>
<ul>
<li>e.g. reading sensors; taking audio samples</li>
<li>can do it continuous</li>
</ul>
<li>Consumer is processing data</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Producer -------&gt; [|||||||||] --------&gt; Consumer</span></div>
</div>
<div>
<ul style="text-align: left;">
<li>Buffer &amp; blocking help equalizing speeds of producer(sender) and consumer(receiver)</li>
<li>if producer is producing data faster than consumer can consume, producer has to block, to slow down producing data - when buffer is full</li>
<li>if consumer is consuming data faster than producer can produce, consumer has to block - when there is no data in a buffer, when buffer is empty</li>
</ul>
<div>
<br></div>
</div>
<div>
<h2 style="text-align: left;">
Week 4 -&nbsp;&nbsp;SYNCHRONIZED COMMUNICATION</h2>
<h3 style="text-align: left;">
Blocking on Channels</h3>
<h4 style="text-align: left;">
Iterating through a Channel</h4>
<br>
<ul style="text-align: left;">
<li>common operation on channel is to iteratively read from channel</li>
<li>this would happen when we have producer and consumer</li>
<ul>
<li>consumer wants to continuously receive data from a channel and process it</li>
</ul>
<li>there is a construct in Go which is made specifically to do this:</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for i:= range c {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Println(i)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<br>
<br>
<ul style="text-align: left;">
<li>continues to read from channel c</li>
<li>one iteration each time a new value is received (is available in channel)</li>
<li>i is assigned to the read value</li>
<li>this <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for </span>loop could be an infinite loop</li>
<ul>
<li>to quit the loop sender can close the channel</li>
<li>loop quits when sender calls <span style="color: blue; font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><b>close(c)</b></span></li>
<li>this is another method that can be performed on a channel</li>
<li>when sender closes channel that gives a signal to the receiver - for loop ends</li>
<li>if we use range to read from channel then we need to call <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">close(c)</span>&nbsp;</li>
</ul>
</ul>
<br><ul style="text-align: left;"><ul>
</ul>
</ul>
<h4 style="text-align: left;">
Receiving from Multiple Goroutines</h4>
<div>
<ul style="text-align: left;">
<li>another common scenario is reading from multiple goroutines or multiple channels that can be associated with multiple goroutines</li>
<li>multiple channels may be used to receive from multiple sources</li>
<li>let's say we have 3 goroutines that are communicating with 2 channels</li>
<ul>
<li>e.g. task3 tries to compute a product of two numbers, each coming from a different channel</li>
</ul>
</ul>
</div>
<div>
<br></div>
<div>
<span style="font-family: Courier New, Courier, monospace;">task1 ----- c1 -----&gt; task3 &lt;----- c2 ------ task2</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>data from both sources might be needed</li>
<li>read sequentially</li>
</ul>
<div>
<span style="font-family: Courier New, Courier, monospace;">a := &lt;- c1</span></div>
</div>
<div>
<span style="font-family: Courier New, Courier, monospace;">b := &lt;- c2</span></div>
<div>
<span style="font-family: Courier New, Courier, monospace;">fmt.Println(a * b)</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>this is blocking</li>
<ul>
<li>T3 first had to wait data to appear on c1 and then for data to appear on c2</li>
</ul>
<li>eventually, T3 will both data and complete task</li>
</ul>
<br><ul style="text-align: left;">
</ul>
</div>
<h4 style="text-align: left;">
Select Statement</h4>
<div>
<ul style="text-align: left;">
<li>sometimes task need data from either channel, from either c1 OR c2,&nbsp;</li>
<li>if we have a choice of multiple channels and want to use the data that comes first ("first come, first served"), no matter which channel from</li>
<ul>
<li>we don't want to read from all channels</li>
<li>we don't want to block (wait on data) on some channel e.g. c1 as it might never happen - in the meantime data might be available on some other channel e.g. c2</li>
<li>we don't know on which channel data will come first</li>
<li>in this case we'll use the select statement:</li>
</ul>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><b><span style="color: blue;">select</span></b> {</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case a = &lt;- c1:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println(a)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case b = &lt;- c2:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println(b)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul>
<li>whichever case happens first, its print will be executed</li>
</ul>
<br><ul>
</ul>
<h4 style="text-align: left;">
Select Send or Receive</h4>
</div>
<div>
<ul style="text-align: left;">
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">select </span>allows choosing data from several channels</li>
<li>we don't have to block on all channels, we practically block only on channel which will first return data</li>
<li>we are blocking here on receiving data but we can also block on sending data</li>
<li>with select, case can be on receiving or sending data:</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">select {</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case a &lt;- inchan:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println(a)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case outchan &lt;- b:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println("sent b")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>if something comes on inchan, a gets that value</li>
<li>we also want to write b value to outchan&nbsp;</li>
<li>inchan is blocked if noone is writing to it</li>
<li>outchan is blocked if noone is reading from it</li>
<li>either of these two actions (cases) can happen first e.g. inchan might have available data before outchan is emptied so can receive new value</li>
<li>whichever thing happens first that case will be executed</li>
<li>if some data comes to inchan before data on outchan becomes available then case 1 will be executed</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Select with Abort Channel</h4>
</div>
<div>
<ul style="text-align: left;">
<li>one common use of select is to have a separate abort channel</li>
<li>producer-consumer scenario</li>
<li>use <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">select</span> with a <b>separate abort channel</b></li>
<li>may want to receive data until an <b>abort signal</b> is received</li>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for { // infinite loop</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;select {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; case a &lt;- c:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;fmt.Println(a) // process data</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; case &lt;-abort:&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;return&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;// abort signal received</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>if anything comes to an abort channel we quit the loop</li>
<li>we don't pay attention which data is coming on abort channel - we dismiss it&nbsp;</li>
</ul>
<br><ul style="text-align: left;">
</ul>
</div>
<h4 style="text-align: left;">
Dafault Select</h4>
<div>
<ul style="text-align: left;">
<li>we have regular cases - e.g. waiting for some channels e.g. c1 and c2</li>
<li>default case: executed if no other cases are satisfied</li>
<li>in this case it will not block at all!</li>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">select {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case a = &lt;- c1:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println(a)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;case b = &lt;- c2:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println(b)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;default:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println("nop")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<h3 style="text-align: left;">
Mutual Exclusion</h3>
<h4 style="text-align: left;">
Goroutines Sharing Variables</h4>
<div>
<ul style="text-align: left;">
<li>sharing variables between goroutines (concurrently) can cause problems&nbsp;</li>
<li>two goroutnes writing to the same shared variable can interfere with each other</li>
<li>function/goroutine is said to be <b>concurrency-safe</b> if can be executed concurrently with other goroutines without interfering improperly with them</li>
<ul>
<li>e.g. it will not alterate variables in other goroutines in some unexpected/unintended/unsafe way</li>
</ul>
</ul>
</div>
<h4 style="text-align: left;">
Variable Sharing Example</h4>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var i int = 0&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var wg sync.WaitGroup</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func inc() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;i = i + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Done()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Add(2)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go inc()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go inc()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Wait()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Println(i)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>two goroutines write to i</li>
<li>i should equal 2</li>
<li>BUT this doesn't always happen</li>
</ul>
<h4 style="text-align: left;">
Possible Interleavings</h4>
</div>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 0</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1: i = i + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2: i = i + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 2</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 0</span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2: i = i + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1: i = i + 1</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 2</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>seems like there is no problem</li>
<li>BUT that is deceiving as there are more interleavings than we think</li>
</ul>
<h4 style="text-align: left;">
Granularity of Concurrency</h4>
</div>
<h4 style="text-align: left;">
<ul style="text-align: left;">
<li><span style="font-weight: normal;">concurrency is at the machine code level, NOT the source code level!</span></li>
<li><span style="font-weight: normal;">intreleavings are not of Go source code instructions, what actually gets interleaved is underlying machine code&nbsp;</span></li>
<li><span style="font-weight: normal;">Go source code is compiled to machine code</span></li>
<li><span style="font-weight: normal;">machine instructions get interleaved</span></li>
<li><span style="font-weight: 400;">interleaving can start in the middle of some Go source code instruction</span></li>
<li><span style="font-weight: 400;">i = i + 1 might be mapped into three machine code instructions:</span></li>
<ul>
<li><span style="font-weight: 400;">read i (read value from memory and place it in the registry)</span></li>
<li><span style="font-weight: 400;">increment (in registry)</span></li>
<li><span style="font-weight: 400;">write i (write it back to memory)</span></li>
</ul>
<li><span style="font-weight: 400;">interleaving happens at this level</span></li>
<li><span style="font-weight: 400;">interleaving machine instructions causes unexpected problems</span></li>
</ul>
</h4>
<h4 style="text-align: left;">
Interleaving Machine Instructions</h4>
<div>
<ul style="text-align: left;">
<li><span style="font-weight: 400;">Both tasks read 0 for i value</span></li>
<li><span style="font-weight: 400;">each task is using its own registry</span></li>
<li><span style="font-weight: 400;">both tasks are sharing variable <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i</span></span></li>
</ul>
</div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i == 0</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1: read i // 0</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2: read i // 0</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1: inc // 1&nbsp;</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task1: write // 1</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i == 1</span></span></div>
<div>
<span style="font-weight: 400;"><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2: inc&nbsp; // 1</span></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Task2: write // overwrites 1 with the same value - 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i == 1</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<h4 style="text-align: left;">
Mutex</h4>
<div>
<ul style="text-align: left;">
<li>how to we do sharing of data correctly between two goroutines?</li>
<li>don't let two goroutines write to a shared variable at the same time&nbsp;</li>
<li>we need to restrict possible interleavings in such way that they don't write to shared variable at the same time</li>
<li>access to shared variables cannot be interleaved</li>
<li>Mutual Exclusion</li>
<ul>
<li>declare code segments in different goroutines which cannot execute concurrently =&gt; they cannot be interleaved</li>
</ul>
<li>writing to shared variables should be mutually exclusive</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Sync.Mutex</h4>
</div>
<div>
<ul style="text-align: left;">
<li>A Mutex ensures mutual exclusion</li>
<li>uses a <b>binary semaphore</b></li>
<ul>
<li>if flag is up</li>
<ul>
<li>shared variable is in use by somebody&nbsp;</li>
<li>when one goroutine is writing into it</li>
<li>only one goroutine can write into variable at a time</li>
<li>once goroutine is using shared variable it has to put the flag up</li>
<li>once goroutine is done with using shared variable it has to put the flag down</li>
</ul>
<li>if flag is down</li>
<ul>
<li>shared variable is available</li>
<li>if another goroutine see that flag is down it knows it can use the shared variable but first it has to put the flag up</li>
</ul>
</ul>
</ul>
<br><ul style="text-align: left;"><ul><ul>
</ul>
</ul>
</ul>
<h4 style="text-align: left;">
Mutex Methods</h4>
</div>
<div>
<ul style="text-align: left;">
<li>putting flag up and down are implemented in methods lock and unlock</li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Lock()</span></li>
<ul>
<li>method puts the flag up (if none of other goroutines has already put the flag up)</li>
<li>notifies others that shared variable is in use</li>
<li>if second goroutine also calls Lock() it will be blocked, it has to wait until first goroutine releases the lock</li>
<li>we can have any number of goroutines (not just two) competing to put the flag up</li>
</ul>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Unlock() </span>method puts the flag down</li>
<ul>
<li>notifies others that it is done with using shared variable</li>
</ul>
<li>When <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Unlock()</span> is called, a blocked <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Lock()</span> can proceed</li>
<li>so in a goroutine we have to put <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Lock()</span> at the beginning of the mutually exclusive region and call <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Unlock()</span> at the end of it; this will ensure that only one goroutine will be in this mutually exclusive region</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Using Sync.Mutex</h4>
</div>
<div>
<ul style="text-align: left;">
<li>Let's fix the code with double increment</li>
<li>Increment operation is now mutually exclusive</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var i int = 0</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var mut sync.Mutex</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func int() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;<b>mut.<span style="color: blue;">Lock()</span></b></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;i = i + 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;<b>mut.<span style="color: blue;">Unlock()</span></b></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>this ensures that reading i from memory, incrementing it and writing the result will be done in one go for each task, context won't be switched in the middle of this Go command line</li>
</ul>
</div>
<div>
<br></div>
<h3 style="text-align: left;">
Once Synchronization</h3>
<div>
<br>
<ul style="text-align: left;">
<li>sync package gives a set of methods that can be used&nbsp; for synchronization between goroutines</li>
</ul>
<br>
<h4 style="text-align: left;">
Synchronous Initialization</h4>
<ul style="text-align: left;">
<li>useful idiom</li>
<li>let's assume we have multi-threaded program and we need to perform initialization</li>
<li>initialization</li>
<ul>
<li>must happen once</li>
<li>must happen before everything else</li>
</ul>
<li>the order of the execution of goroutines is not known so how can we determine which goroutine shall perform the initialization?</li>
<li>How do you&nbsp; perform initialization with multiple goroutines?</li>
<li>Could perform initialization before starting the goroutines</li>
<ul>
<li>put it at the beginning of main() - the first goroutine</li>
<li>sometimes this might not be an option</li>
</ul>
<li>another option is using once.Do()</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Sync.Once</h4>
<div>
<ul style="text-align: left;">
<li>has one method, <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><b><span style="color: blue;">once.Do</span></b>(f)</span></li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">once.Do(f)</span> can be put (called) in many goroutines but function <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">f</span>&nbsp;(e.g. some initialization) will be <u>executed only once</u></li>
<li>all calls to once.Do(f) block until the first returns&nbsp;</li>
<ul>
<li>this ensures that <u>initialization is executed first</u></li>
</ul>
</ul>
<u><br></u><ul style="text-align: left;"><ul>
</ul>
</ul>
<h4 style="text-align: left;">
Sync.Once Example</h4>
</div>
<div>
<ul style="text-align: left;">
<li>make two goroutines, initialization only once</li>
<li>each goroutine executes <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">doStuff()</span></li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">doStuff()</span> performs initialization at the beginning of it</li>
<li>only one goroutine has to perform initialization within <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">doStuff()</span></li>
<li><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">setup()</span> should execute only once</li>
<ul>
<li>only one goroutine will executed it, although it is called in both&nbsp;</li>
</ul>
<li>"hello" should not print until <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">setup()</span> returns</li>
</ul>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var wg sync.WaitGroup</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">var on sync.Once</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func Setup() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Println("Init")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func doStuff() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;on.Do(Setup)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;fmt.Println("hello")</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Done()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Add(2)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go doStuff()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go doStuff()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Wait()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Output:</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">Init // result of Setup()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">hello // from one goroutine&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">hello // from another goroutine&nbsp;</span></div>
<div>
<br></div>
<h4 style="text-align: left;">
Deadlock</h4>
<div>
<ul style="text-align: left;">
<li>it should be avoided when coding</li>
<li>comes from synchronization dependencies</li>
<li>if we have multiple goroutines and synchronization can cause one goroutine's execution to depend on another's</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Synchronization Dependencies</h4>
</div>
<div>
<ul style="text-align: left;">
<li>synchronization causes the execution of different goroutines (e.g. G1 and G2) to depend on each other</li>
</ul>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">G1:</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">ch &lt;- 1</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">mut.Unlock()</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">G2:</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">x := &lt;- ch&nbsp;</span><br>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">mut.Lock()</span><br>
<br>
<ul style="text-align: left;">
<li>there is a blocking dependencies here:&nbsp;</li>
<ul>
<li>G1 writes to channel and G2 reads from it so G2 is dependent on G1 as G2 blocks until G1 is executed; G2 cannot continue until G1 does something</li>
<li>G2 can acquire a lock only after G1 releases it</li>
</ul>
</ul>
<br><ul style="text-align: left;"><ul>
</ul>
</ul>
<h4 style="text-align: left;">
Deadlock</h4>
<div>
<ul style="text-align: left;">
<li>circular dependencies cause all involved goroutines to block&nbsp;</li>
<ul>
<li>G1 waits for G2</li>
<li>G2 waits for G1</li>
</ul>
<li>can be caused by waiting on channels too</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Deadlock Example</h4>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func dostuff(c1 chan int, c2 chan int) {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;&lt;- c1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;c2 &lt;- 1</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Done()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>read from first channel</li>
<ul>
<li>wait for write onto first channel</li>
</ul>
<li>write to second channel</li>
<ul>
<li>wait for read from second channel</li>
</ul>
</ul>
<div>
<br></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func main() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;ch1 := make(chan int)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;ch2 := make(chan int)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Add()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go doStuff(ch1, ch2)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go doStuff(ch2, ch1)</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;wg.Wait()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>doStuff() argument order is swapped</li>
<li>Each goroutine blocked on a channel read</li>
<li>nothing can progress =&gt; deadlock!</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Deadlock Detection</h4>
</div>
<div>
<ul style="text-align: left;">
<li>&nbsp;Golang runtime automatically detects when all goroutines are deadlocked:</li>
<ul>
<li>"fatal error: all goroutines are asleep - deadlock!"</li>
</ul>
<li>however, it cannot detect when a subset of goroutines are deadlocked</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Dining Philosophers Problem</h4>
</div>
<div>
<ul style="text-align: left;">
<li>classic concurrency problem, with a deadlock</li>
<li>Problem:</li>
<ul>
<li>5 philosophers sitting at a round table</li>
<li>each one has a plate of rice</li>
<li>1 chopstick is placed between each adjacent pair</li>
<li>to eat, a philosopher need two chopsticks, one from the left and one from the right of his plate</li>
<li>only one philosopher can hold a chopstick at a time</li>
<li>not enough chopsticks for everyone to eat at once</li>
</ul>
</ul>
<br><ul style="text-align: left;"><ul>
</ul>
</ul>
<h4 style="text-align: left;">
Dining Philosophers Issues</h4>
</div>
<div>
<br></div>
<div>
&nbsp; &nbsp; \ O | O /&nbsp;</div>
<div>
O&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;O</div>
<div>
&nbsp; &nbsp; /&nbsp; &nbsp; O&nbsp; &nbsp; \</div>
<div>
<ul style="text-align: left;">
<li>each chopstick is a mutex because philosophers have mutually exclusive access to it</li>
<li>each philosopher is associated with a goroutine and two chopsticks</li>
</ul>
<br><ul style="text-align: left;">
</ul>
<h4 style="text-align: left;">
Chopsticks and Philosophers</h4>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">type ChopS structure {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;sync.Mutex</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">type Philo struct {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;leftCS, rightCS *ChopS</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<h4 style="text-align: left;">
Philosopher Eat Method</h4>
<div>
<br></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">func (p Philo) eat() {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;for {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; p.leftCS.Lock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; p.rigthCS.Lock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; fmt.Println("eating")</span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; p.rigthCS.Unlock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; p.leftCS.Unlock()</span></div>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<h4 style="text-align: left;">
Initialization in Main</h4>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">CSticks := make([]*ChopS, 5) // create slice</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for i:=0; i&lt;5; i++ {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;CSticks[i] = new(ChopS) // BK: why not using &amp;ChopS?</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">philos := make([]*Philo, 5)&nbsp;</span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">// create slice</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for i:=0; i&lt;5; i++ {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;philos[i] = &amp;Philo{</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; CSticks[i],&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp; &nbsp; CSticks[(i + 1)%5]</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;}</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
</div>
<div>
<ul style="text-align: left;">
<li>Initialize chopsticks and philosophers</li>
<li>Notice <span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">(i + 1)%5</span></li>
<ul>
<li><span style="font-family: inherit;">this is because we can't use simply</span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"> i + 1</span><span style="font-family: inherit;"> as for </span><span style="font-family: &quot;arial&quot; , &quot;helvetica&quot; , sans-serif;">i = 4</span><span style="font-family: inherit;">, we'd have </span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">i = 5</span><span style="font-family: inherit;"> but that should be 0 (</span><span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">5 % 5 = 0</span><span style="font-family: inherit;">)</span></li>
</ul>
</ul>
</div>
<h4 style="text-align: left;">
Start the dining in Main</h4>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;"><br></span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">for i := 0; i &lt; 5; i++ {</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;go philos[i].eat()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<br></div>
<div>
<ul style="text-align: left;">
<li>start each philosopher eating</li>
<li>would also need Wait in main() so main() does not return&nbsp; before philosophers complete eating</li>
<li>that was the code written in naive way; naive for we'd have a deadlock!</li>
</ul>
<h4 style="text-align: left;">
Deadlock Problem</h4>
</div>
<div>
<br></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">p.leftCS.Lock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">p.rigthCS.Lock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">fmt.Println("eating")</span></div>
<div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">p.rigthCS.Unlock()</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">p.leftCS.Unlock()</span></div>
</div>
</div>
<div>
<br></div>
</div>
<div>
<ul style="text-align: left;">
<li>this sequence is executed in each of 5 goroutines</li>
<li>these goroutines are ordered in non-deterministic way</li>
<li>in one possible interleaving all philosophers might lock their left chopsticks concurrently</li>
<li>all chopsticks would be locked =&gt; none can lock their right chopsticks cause someone's right chopstick is someone else's left one and is locked</li>
<li>in such interleaving we are in a deadlock!</li>
</ul>
<h4 style="text-align: left;">
Deadlock Solution</h4>
</div>
<div>
<ul style="text-align: left;">
<li>Dijkstra's solution: each philosopher picks up lowest numbered chopstick first</li>
<li>our code is NOT doing that:</li>
</ul>
</div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">philos[i] = &amp;Philo{</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;CSticks[i],&nbsp;</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">&nbsp; &nbsp;CSticks[(i+1)%5]</span></div>
<div>
<span style="font-family: &quot;courier new&quot; , &quot;courier&quot; , monospace;">}</span></div>
<div>
<ul style="text-align: left;">
<li>In our current code:</li>
<ul>
<li>Philosopher 0 picks first CS0 as left and CS1 as right</li>
<ul>
<li>this is OK as it picked CS with lowest number first</li>
</ul>
<li>Philosopher 1 picks first CS1 as left and CS2 as right</li>
<ul>
<li>this is OK as it picked CS with lowest number first</li>
</ul>
<li>Philosopher 2 picks first CS2 as left and CS3 as right</li>
<ul>
<li>this is OK as it picked CS with lowest number first</li>
</ul>
<li>Philosopher 3 picks first CS3 as left and CS4 as right</li>
<ul>
<li>this is OK as it picked CS with lowest number first</li>
</ul>
<li>Philosopher 4 picks first CS4 as left and CS0 as right</li>
<ul>
<li>this is <u>NOT</u> OK as it picked CS with higher number first</li>
</ul>
</ul>
<li>Philosopher 4 picks up chopstick 4 before chopstick 0</li>
<ul>
<li>this violates Dijkstra's law</li>
</ul>
<li>With Dijkstra's solution:</li>
<ul>
<li>in one interleaving we may have:</li>
<ul>
<li>Philosopher 0 picks first CS0 as left and CS1 as right</li>
<li>Philosopher 1 waits for PH0 to finish so can grab CS1 first; once PH0 finishes eating: Philosopher 1 picks first CS1 as left and CS2 as right</li>
<li>PH2 waits for PH1 to finish so can grab CS2 first; once PH1 finishes eating: PH2 picks first CS2 as left and CS3 as right</li>
<li>PH3 waits for PH2 to finish so can grab CS3 first; once PH2 finishes eating: PH3 picks first CS3 as left and CS4 as right</li>
<li>PH4 waits for PH3 to finish so can grab CS4 first; once PH3 finishes eating: PH4 picks first CS4 as left and CS0 as right</li>
</ul>
<li>Philospher 4 would first pick CS0 and then CS4 BUT Philosopher 0 blocks allowing Philosopher 4 to eat as PH0 has already picked SC0 so PH4 can't pick it; PH4 is blocked to pick CS0 so it won't pick CS4 either allowing Philosopher3 to pick both CS</li>
<li>No deadlock, but philosopher 4 may starve</li>
<li>Philospher 4 gets the lowest priority most of the time; he has to wait for the other most of the time</li>
<li>starvation: some goroutines may not be scheduled to be executed as often as others</li>
<li>this happens when we have circular dependency</li>
<li><b>starvation</b> is not ideal but <b>deadlock</b> is the worst</li>
</ul>
</ul>
</div>
