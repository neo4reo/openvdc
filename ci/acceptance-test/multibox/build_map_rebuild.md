<code>&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L45">1-Put&nbsp;default&nbsp;config&nbsp;file&nbsp;in&nbsp;place&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L35">2-Download&nbsp;minimal-7.2.1511-x86_64.kvm.box&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L9">3-Create&nbsp;bridge&nbsp;vdc_env_br0&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L30">4.0-Cleanup&nbsp;old&nbsp;environment&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L35">4.1.0-Destroying&nbsp;10.0.100.10-zookeeper&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L29">4.1.1-Remove&nbsp;SSH&nbsp;key&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L19">4.1.2-Kill&nbsp;vm&nbsp;zookeeper&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L36">4.1.3-Remove&nbsp;vm&nbsp;zookeeper&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L35">4.2.0-Destroying&nbsp;10.0.100.11-mesos-master&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L29">4.2.1-Remove&nbsp;SSH&nbsp;key&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L19">4.2.2-Kill&nbsp;vm&nbsp;mesos&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L36">4.2.3-Remove&nbsp;vm&nbsp;mesos&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L35">4.3.0-Destroying&nbsp;10.0.100.12-vdc-scheduler&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L29">4.3.1-Remove&nbsp;SSH&nbsp;key&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L19">4.3.2-Kill&nbsp;vm&nbsp;scheduler&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L36">4.3.3-Remove&nbsp;vm&nbsp;scheduler&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L35">4.4.0-Destroying&nbsp;10.0.100.13-vdc-executor-null&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L29">4.4.1-Remove&nbsp;SSH&nbsp;key&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L19">4.4.2-Kill&nbsp;vm&nbsp;executor-null&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L36">4.4.3-Remove&nbsp;vm&nbsp;executor-null&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L35">4.5.0-Destroying&nbsp;10.0.100.14-vdc-executor-lxc&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L29">4.5.1-Remove&nbsp;SSH&nbsp;key&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L19">4.5.2-Kill&nbsp;vm&nbsp;executor-lxc&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L36">4.5.3-Remove&nbsp;vm&nbsp;executor-lxc&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L43">4.6-Create&nbsp;cache&nbsp;folder&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L60">5.0-Building&nbsp;10.0.100.10-zookeeper&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L7">5.1-Deploy&nbsp;seed&nbsp;image&nbsp;for&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L15">5.2-Mount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L27">5.3-Synching&nbsp;guestroot&nbsp;for&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesosphere/init.sh#L3">5.4-Download&nbsp;mesosphere&nbsp;repo&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L4">5.5-Create&nbsp;the&nbsp;public&nbsp;key&nbsp;and&nbsp;setup&nbsp;ssh&nbsp;config&nbsp;for&nbsp;&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L13">5.6-Install&nbsp;authorized&nbsp;ssh&nbsp;key&nbsp;for&nbsp;&nbsp;on&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-zookeeper/install.sh#L2">5.7-Install&nbsp;zookeeper&nbsp;on&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L4">5.8.0-Cleanup&nbsp;build&nbsp;phase&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L47">5.8.1-Unmount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;zookeeper&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L9">5.8.2-Cache&nbsp;box&nbsp;for&nbsp;branch&nbsp;develop&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L16">5.8.3-Delete&nbsp;raw&nbsp;image&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L24">5.9-Start&nbsp;kvm&nbsp;for&nbsp;zookeeper&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L56">5.10-Attach&nbsp;zookeeper_tap&nbsp;to&nbsp;vdc_env_br0&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L4">5.11-Import&nbsp;ssh&nbsp;key&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L12">5.12-Wait&nbsp;for&nbsp;ssh&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L60">6.0-Building&nbsp;10.0.100.11-mesos-master&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L7">6.1-Deploy&nbsp;seed&nbsp;image&nbsp;for&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L15">6.2-Mount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L27">6.3-Synching&nbsp;guestroot&nbsp;for&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesosphere/init.sh#L3">6.4-Download&nbsp;mesosphere&nbsp;repo&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L4">6.5-Create&nbsp;the&nbsp;public&nbsp;key&nbsp;and&nbsp;setup&nbsp;ssh&nbsp;config&nbsp;for&nbsp;&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L13">6.6-Install&nbsp;authorized&nbsp;ssh&nbsp;key&nbsp;for&nbsp;&nbsp;on&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesos/install.sh#L3">6.7-Install&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L4">6.8.0-Cleanup&nbsp;build&nbsp;phase&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L47">6.8.1-Unmount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;mesos&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L9">6.8.2-Cache&nbsp;box&nbsp;for&nbsp;branch&nbsp;develop&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L16">6.8.3-Delete&nbsp;raw&nbsp;image&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L24">6.9-Start&nbsp;kvm&nbsp;for&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L56">6.10-Attach&nbsp;mesos&nbsp;to&nbsp;vdc_env_br0&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L4">6.11-Import&nbsp;ssh&nbsp;key&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L12">6.12-Wait&nbsp;for&nbsp;ssh&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">6.13-Disable&nbsp;service:&nbsp;mesos-slave&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L60">7.0-Building&nbsp;10.0.100.12-vdc-scheduler&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L7">7.1-Deploy&nbsp;seed&nbsp;image&nbsp;for&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L15">7.2-Mount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L27">7.3-Synching&nbsp;guestroot&nbsp;for&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesosphere/init.sh#L3">7.4-Download&nbsp;mesosphere&nbsp;repo&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L4">7.5-Create&nbsp;the&nbsp;public&nbsp;key&nbsp;and&nbsp;setup&nbsp;ssh&nbsp;config&nbsp;for&nbsp;&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L13">7.6-Install&nbsp;authorized&nbsp;ssh&nbsp;key&nbsp;for&nbsp;&nbsp;on&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-epel/install.sh#L4">7.7-Install&nbsp;EPEL&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/install.sh#L4">7.8-Install&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesos/install.sh#L3">7.9-Install&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-zookeeper/install.sh#L2">7.10-Install&nbsp;zookeeper&nbsp;on&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L4">7.11.0-Cleanup&nbsp;build&nbsp;phase&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L47">7.11.1-Unmount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;scheduler&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L9">7.11.2-Cache&nbsp;box&nbsp;for&nbsp;branch&nbsp;develop&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L16">7.11.3-Delete&nbsp;raw&nbsp;image&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L24">7.12-Start&nbsp;kvm&nbsp;for&nbsp;scheduler&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L56">7.13-Attach&nbsp;vdc_scheduler_tap&nbsp;to&nbsp;vdc_env_br0&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L4">7.14-Import&nbsp;ssh&nbsp;key&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L12">7.15-Wait&nbsp;for&nbsp;ssh&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/postconfigure.sh#L4">7.16-Setup&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">7.17-Disable&nbsp;service:&nbsp;mesos-master&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">7.18-Disable&nbsp;service:&nbsp;mesos-slave&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">7.19-Disable&nbsp;service:&nbsp;zookeeper&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L60">8.0-Building&nbsp;10.0.100.13-vdc-executor-null&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L7">8.1-Deploy&nbsp;seed&nbsp;image&nbsp;for&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L15">8.2-Mount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L27">8.3-Synching&nbsp;guestroot&nbsp;for&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesosphere/init.sh#L3">8.4-Download&nbsp;mesosphere&nbsp;repo&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L4">8.5-Create&nbsp;the&nbsp;public&nbsp;key&nbsp;and&nbsp;setup&nbsp;ssh&nbsp;config&nbsp;for&nbsp;&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L13">8.6-Install&nbsp;authorized&nbsp;ssh&nbsp;key&nbsp;for&nbsp;&nbsp;on&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-epel/install.sh#L4">8.7-Install&nbsp;EPEL&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/install.sh#L4">8.8-Install&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesos/install.sh#L3">8.9-Install&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-zookeeper/install.sh#L2">8.10-Install&nbsp;zookeeper&nbsp;on&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L4">8.11.0-Cleanup&nbsp;build&nbsp;phase&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L47">8.11.1-Unmount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;executor-null&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L9">8.11.2-Cache&nbsp;box&nbsp;for&nbsp;branch&nbsp;develop&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L16">8.11.3-Delete&nbsp;raw&nbsp;image&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L24">8.12-Start&nbsp;kvm&nbsp;for&nbsp;executor-null&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L56">8.13-Attach&nbsp;vdc_executor_null_tap&nbsp;to&nbsp;vdc_env_br0&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L4">8.14-Import&nbsp;ssh&nbsp;key&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L12">8.15-Wait&nbsp;for&nbsp;ssh&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/postconfigure.sh#L4">8.16-Setup&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">8.17-Disable&nbsp;service:&nbsp;mesos-master&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">8.18-Disable&nbsp;service:&nbsp;zookeeper&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./build.sh#L60">9.0-Building&nbsp;10.0.100.14-vdc-executor-lxc&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L7">9.1-Deploy&nbsp;seed&nbsp;image&nbsp;for&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L15">9.2-Mount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/init.sh#L27">9.3-Synching&nbsp;guestroot&nbsp;for&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesosphere/init.sh#L3">9.4-Download&nbsp;mesosphere&nbsp;repo&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L4">9.5-Create&nbsp;the&nbsp;public&nbsp;key&nbsp;and&nbsp;setup&nbsp;ssh&nbsp;config&nbsp;for&nbsp;&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/install.sh#L13">9.6-Install&nbsp;authorized&nbsp;ssh&nbsp;key&nbsp;for&nbsp;&nbsp;on&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-epel/install.sh#L4">9.7-Install&nbsp;EPEL&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/install.sh#L4">9.8-Install&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-mesos/install.sh#L3">9.9-Install&nbsp;mesos&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-zookeeper/install.sh#L2">9.10-Install&nbsp;zookeeper&nbsp;on&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L4">9.11.0-Cleanup&nbsp;build&nbsp;phase&nbsp;:::</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/common.source#L47">9.11.1-Unmount&nbsp;temporary&nbsp;root&nbsp;folder&nbsp;for&nbsp;executor-lxc&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L9">9.11.2-Cache&nbsp;box&nbsp;for&nbsp;branch&nbsp;develop&nbsp;(not&nbsp;done,skippable)</a></code><br>
<code>&#42;&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L16">9.11.3-Delete&nbsp;raw&nbsp;image&nbsp;(DONE,skippable)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L24">9.12-Start&nbsp;kvm&nbsp;for&nbsp;executor-lxc&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-box/boot.sh#L56">9.13-Attach&nbsp;vdc_executor_lxc_tap&nbsp;to&nbsp;vdc_env_br0&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L4">9.14-Import&nbsp;ssh&nbsp;key&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-ssh/boot.sh#L12">9.15-Wait&nbsp;for&nbsp;ssh&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/step-lxc/postconfigure.sh#L4">9.16-Setup&nbsp;LXC&nbsp;(not&nbsp;done)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">9.17-Disable&nbsp;service:&nbsp;mesos-master&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L113">9.18-Disable&nbsp;service:&nbsp;zookeeper&nbsp;(DONE)</a></code><br>
<code>&#42;&#42;&nbsp;&nbsp;<a href="/./ind-steps/common.source#L56">10-Masquerade&nbsp;for&nbsp;subnet&nbsp;10.0.100.0/24&nbsp;(DONE)</a></code><br>