public class Twofer {
    public String twofer(String name) {
        return name == null || name.isEmpty() ? "One for you, one for me."
                : String.format("One for %s, one for me.", name);
    }
}
