import { useState } from 'react';

const useWindowDimension = () => {

    const [dimen,setdimen] = useState({
        woh: window.outerHeight,
        wow: window.outerWidth
    })

    return dimen;
}

export default useWindowDimension;