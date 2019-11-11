mod util {
   pub fn add(a:i32, b:i32) -> i32 {
       a+b
    }
}
#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn it_works() {
        assert_eq!(util::add(1,1), 2);
    }
    #[test]
    fn it_will_works() {
        assert_eq!(util::add(1,1), 3);
    }
}
