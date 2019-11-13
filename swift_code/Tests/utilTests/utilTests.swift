import XCTest
@testable import util

final class utilTests: XCTestCase {
    func testExample() {
        /
        XCTAssertEqual(Util.add(1,2), 2)
    }

    static var allTests = [
        ("testExample", testExample),
    ]
}
