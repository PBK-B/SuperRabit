const useSystemPlatform = () => {
    var u = navigator.userAgent;
    if( u.indexOf('Android') > -1 || u.indexOf('Linux') > -1){
        return 'android'
    }else if(u.indexOf('iPhone') > -1){
        return 'ios'
    }else if(u.indexOf('iPad') > -1){
        return 'ipad'
    }else if(u.indexOf('Windows NT') > -1){
        return 'windows'
    }else if(u.indexOf('Mac') > -1){
        return 'mac'
    }else{
        return 'unknown'
    }
}
export default useSystemPlatform;