package solace.coverage.paths;

import static org.junit.Assert.assertEquals;

import java.util.HashMap;
import java.util.Map;

import com.tngtech.java.junit.dataprovider.DataProvider;
import com.tngtech.java.junit.dataprovider.DataProviderRunner;
import com.tngtech.java.junit.dataprovider.UseDataProvider;
import org.junit.Test;
import org.junit.runner.RunWith;

@RunWith(DataProviderRunner.class)
public class PathsTreeTest {

    @Test
    @UseDataProvider("pathTreeCreateData")
    public void CreatePathTreeTest(String[] paths, Integer[] values) {
        Map<String, Integer> pathsMap = buildMapFromArrays(paths, values);
        PathsTree<Integer> pathsTree = new PathsTree<>(pathsMap);
        for (Integer i = 0; i < paths.length; i++)
            assertEquals(values[i], pathsTree.resolve(paths[i])); 
    }

    @DataProvider
    public static Object[][] pathTreeCreateData() {
        return new Object[][] {
            new Object[] {
                new String[] {
                    "/",
                    "/foo",
                    "/foo/bar",
                    "/foo/bar/baz/foo"
                },
                new Integer[] {
                    1,
                    2,
                    3,
                    4
                }
            },
        };
    }

    @Test
    @UseDataProvider("pathTreeResolveData")
    public void PathTreeResolveWithPathVariables(String[] paths, String[] keys, Integer[] values) {
        Map<String, Integer> pathsMap = buildMapFromArrays(paths, values);
        PathsTree<Integer> pathsTree = new PathsTree<>(pathsMap);
        for (Integer i = 0; i < paths.length; i++)
            assertEquals(values[i], pathsTree.resolve(keys[i])); 
    }
    
    @DataProvider
    public static Object[][] pathTreeResolveData() {
        return new Object[][] {
            new Object[] {
                new String[] {
                    "/foo",
                    "/foo/{bar}",
                    "/foo/{bar}/baz",
                    "/foo/{bar}/{baz}/foo",
                    "/foo/bar/baz"
                },
                new String[] {
                    "/foo",
                    "/foo/123",
                    "/foo/321/baz",
                    "/foo/asdkjdsf/sdlkjsdf/foo",
                    "/foo/bar/baz"
                },
                new Integer[] {
                    1,
                    2,
                    3,
                    4,
                    5
                }
            },
        };
    }

    @Test
    @UseDataProvider("pathTreeResolvePartlyData")
    public void PathTreeResolveWithPathVariablesPartly(String[] paths, String[] keys, Integer[] values) {
        Map<String, Integer> pathsMap = buildMapFromArrays(paths, values);
        PathsTree<Integer> pathsTree = new PathsTree<>(pathsMap);
        for (Integer i = 0; i < paths.length; i++)
            assertEquals(values[i], pathsTree.resolve(keys[i])); 
    }
    
    @DataProvider
    public static Object[][] pathTreeResolvePartlyData() {
        return new Object[][] {
            new Object[] {
                new String[] {
                    "/foo",
                    "/foo/{bar}",
                    "/foo/{bar}:abc/baz",
                    "/foo/err{bar}/{baz}/foo",
                    "/baz/00{aa}00/foo",
                    "/bar/[{ff}]",
                    "/foo/bar/baz"
                },
                new String[] {
                    "/foo",
                    "/foo/123",
                    "/foo/321:abc/baz",
                    "/foo/errasdkjdsf/sdlkjsdf/foo",
                    "/baz/00100/foo",
                    "/bar/[sjdfhsjdf]",
                    "/foo/bar/baz"
                },
                new Integer[] {
                    1,
                    2,
                    3,
                    4,
                    5,
                    6,
                    7
                }
            },
        };
    }

    public static <T, K> Map<T, K> buildMapFromArrays(T[] keys, K[] values) {
        Map<T, K> map = new HashMap<T, K>();
        if (keys.length != values.length)
            throw new IllegalArgumentException("Arrays must be the same length");
        for (Integer i = 0; i < keys.length; i++)
            map.put(keys[i], values[i]);
        return map;
    }
}