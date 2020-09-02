use chrono::{DateTime, Duration, Utc};

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    let t = start.checked_add_signed(Duration::seconds(1000000000));
    return t.unwrap();
}
