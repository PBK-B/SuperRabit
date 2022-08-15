import { useState } from 'react';

const useWindowInnerDimension = () => {

    const [dimen,setdimen] = useState({
        wih: window.innerHeight,
        wiw: window.innerWidth
    })

    return dimen;
}

export default useWindowInnerDimension;