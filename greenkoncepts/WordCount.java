import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Collectors;
import java.util.stream.Stream;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

public class WordCount {

	public static void main(String args[]) {

		String fileName = args[0];
		Map<String, Integer> wordCount = new ConcurrentHashMap<>();

		try (Stream<String> stream = Files.lines(Paths.get(fileName))) {

			//1. filter length >= 2
			//2. convert all content to lower case and counting
			wordCount = stream.parallel().filter(word -> (word.length() >= 2))
					.collect(Collectors.toConcurrentMap(w -> w.toLowerCase(), w -> 1, Integer::sum));

		} catch (IOException e) {
			e.printStackTrace();
		}

		AtomicInteger totalWord = new AtomicInteger(0);
		wordCount.forEach((k, v) -> {
			totalWord.getAndAdd(v);
			System.out.printf("%s, %d\n", k, v);
		});
		System.out.printf("Total words: %d\n", totalWord.intValue());
	}

}
