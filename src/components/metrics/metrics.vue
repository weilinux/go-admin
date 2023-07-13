<template>
  <h4>主机监控</h4>

  <div class="d-flex justify-content-center align-items-center container ">

    主机IP: {{ .host.HostIP }}
    Platformat: {{ .hostInfo.Platform }}
    Host ID: {{ .hostInfo.HostID }}
    Operating System： {{ .hostInfo.OS }}

    Host Name {{ .hostInfo.Hostname }}
    Boot Time {{ .hostInfo.BootTime }}


    % of N CPU cores

    GiB Memory


    MiB/s Disk I/O

    Kbps Network Trffic


</template>

<script>
export default {
  name: "metrics"
}
</script>

<style scoped>

</style>