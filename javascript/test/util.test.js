const {add} = require("../")
describe("Util.js", () => {
    it("should able to add", () =>{
        expect(add(1,1)).toBe(2)
    })
    it("should fail to add", () =>{
        expect(add(1,1)).toBe(3)
    })
})