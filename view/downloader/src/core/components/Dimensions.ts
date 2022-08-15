
type DimenensionType = "window" | "screen"

class Dimensions {

    static get(type: DimenensionType) {
        if(type === "window"){
            return {width: window.outerWidth,height: window.outerHeight}
        }else {
            return {width: window.innerWidth,height: window.innerHeight}
        }
    }
    static addListener(listener: (width:number,height:number) => void) {
        // window.addEventListener('resize',() => {  
        // })
        // window.removeEventListener('resize',(ev:any) => {})
    }

}

export default Dimensions;