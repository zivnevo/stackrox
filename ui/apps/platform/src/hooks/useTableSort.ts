import { useEffect, useState, ReactElement } from 'react';

export type TableColumn = {
    Header: string;
    accessor: string;
    Cell?: ({ original, value }) => ReactElement | string;
    sortField?: string;
};

export type SortDirection = 'asc' | 'desc';
type SortOption = {
    field: string;
    reversed: boolean;
};

export type UseTableSort = {
    activeSortIndex: number;
    setActiveSortIndex: (idx) => void;
    activeSortDirection: SortDirection;
    setActiveSortDirection: (dir) => void;
    sortOption: SortOption;
};

function useTableSort(columns: TableColumn[], defaultSort: SortOption): UseTableSort {
    // index of the currently active column
    const [activeSortIndex, setActiveSortIndex] = useState(0);
    // sort direction of the currently active column
    const [activeSortDirection, setActiveSortDirection] = useState<SortDirection>('desc');

    const [sortOption, setSortOption] = useState<SortOption>(defaultSort);

    useEffect(() => {
        const { sortField } = columns[activeSortIndex];
        if (sortField) {
            const newSortOption: SortOption = {
                field: sortField,
                reversed: activeSortDirection === 'desc',
            };
            setSortOption(newSortOption);
        }
    }, [activeSortIndex, activeSortDirection, columns]);

    return {
        activeSortIndex,
        setActiveSortIndex,
        activeSortDirection,
        setActiveSortDirection,
        sortOption,
    };
}

export default useTableSort;
